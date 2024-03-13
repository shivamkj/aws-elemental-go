package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

func CreateMediaPackageChannelGroup(id string) string {
	channelGroupName := fmt.Sprintf("ChannelGroup-%s", id)
	inputParam := mediapackagev2.CreateChannelGroupInput{
		ChannelGroupName: &channelGroupName,
		Tags:             commonTag,
	}

	output, err := mediaPackageClient.CreateChannelGroup(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}

	return *output.Arn
}
