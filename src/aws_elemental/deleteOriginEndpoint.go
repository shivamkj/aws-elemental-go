package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

func DeleteOriginEndpoint(id string) {
	var (
		channelName      = fmt.Sprintf("Channel-%s", id)
		channelGroupName = fmt.Sprintf("ChannelGroup-%s", id)
		endpointName     = fmt.Sprintf("Endpoint-%s", id)
	)
	inputParam := mediapackagev2.DeleteOriginEndpointInput{
		ChannelName:        &channelName,
		ChannelGroupName:   &channelGroupName,
		OriginEndpointName: &endpointName,
	}

	_, err := mediaPackageClient.DeleteOriginEndpoint(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
