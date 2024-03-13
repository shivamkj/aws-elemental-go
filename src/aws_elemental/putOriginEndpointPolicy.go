package aws_elemental

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

const policyTemplate = `{
	"Version": "2012-10-17",
	"Id": "AnonymousAccessPolicy",
	"Statement": [
		{
			"Sid": "AllowAnonymousAccess",
			"Effect": "Allow",
			"Principal": "*",
			"Action": "mediapackagev2:GetObject",
			"Resource": "arn:aws:mediapackagev2:%s:%s:channelGroup/ChannelGroup-%s/channel/Channel-%s/originEndpoint/Endpoint-%s"
		}
	]
}`

func PutOriginEndpointPolicy(id string) {
	var (
		channelName      = fmt.Sprintf("Channel-%s", id)
		channelGroupName = fmt.Sprintf("ChannelGroup-%s", id)
		endpointName     = fmt.Sprintf("Endpoint-%s", id)
		policy           = fmt.Sprintf(policyTemplate, region, os.Getenv("ENV_ACCOUNT_ID"), id, id, id)
	)
	inputParam := mediapackagev2.PutOriginEndpointPolicyInput{
		ChannelName:        &channelName,
		ChannelGroupName:   &channelGroupName,
		OriginEndpointName: &endpointName,
		Policy:             &policy,
	}

	_, err := mediaPackageClient.PutOriginEndpointPolicy(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
