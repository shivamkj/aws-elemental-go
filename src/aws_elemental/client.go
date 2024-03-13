package aws_elemental

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
	"github.com/joho/godotenv"
)

var (
	mediaLiveClient    *medialive.Client
	mediaPackageClient *mediapackagev2.Client
	region             string
	roleArn            string
	segmentDuration    int32 = 4
	commonTag                = map[string]string{
		"creation": "created by API",
	}
)

func getAwsConfig() aws.Config {
	awscreds := aws.NewCredentialsCache(
		credentials.NewStaticCredentialsProvider(os.Getenv("ENV_ACCESS_KEY_ID"), os.Getenv("ENV_SECRET_ACCESS_KEY"), ""),
	)

	config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(awscreds),
	)
	if err != nil {
		panic(err)
	}

	return config
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("error loading env file")
	}
	region = os.Getenv("ENV_AWS_REGION")
	roleArn = os.Getenv("ENV_ROLE_ARN")

	config := getAwsConfig()

	// Create MediaLive client
	mediaLiveClient = medialive.NewFromConfig(config)

	// Create MediaPackage client
	mediaPackageClient = mediapackagev2.NewFromConfig(config)
}
