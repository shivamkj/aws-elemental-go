package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
)

func DeleteMediaPackageChannelGroup(id string) {
	channelGroupName := fmt.Sprintf("ChannelGroup-%s", id)
	inputParam := mediapackagev2.DeleteChannelGroupInput{
		ChannelGroupName: &channelGroupName,
	}

	_, err := mediaPackageClient.DeleteChannelGroup(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
