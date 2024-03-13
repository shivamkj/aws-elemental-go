package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

func CreateMediaPackageChannel(id string) []string {
	var (
		channelName      = fmt.Sprintf("Channel-%s", id)
		channelGroupName = fmt.Sprintf("ChannelGroup-%s", id)
	)
	inputParam := mediapackagev2.CreateChannelInput{
		ChannelName:      &channelName,
		ChannelGroupName: &channelGroupName,
		Tags:             commonTag,
	}

	output, err := mediaPackageClient.CreateChannel(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}

	return []string{
		*output.IngestEndpoints[0].Url,
		*output.IngestEndpoints[1].Url,
	}
}
