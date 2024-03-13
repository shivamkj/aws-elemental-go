package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

func DeleteMediaPackageChannel(id string) {
	var (
		channelName      = fmt.Sprintf("Channel-%s", id)
		channelGroupName = fmt.Sprintf("ChannelGroup-%s", id)
	)
	inputParam := mediapackagev2.DeleteChannelInput{
		ChannelName:      &channelName,
		ChannelGroupName: &channelGroupName,
	}

	_, err := mediaPackageClient.DeleteChannel(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
