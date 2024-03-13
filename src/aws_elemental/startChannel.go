package aws_elemental

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func StartChannel(channelId string) {
	inputParam := medialive.StartChannelInput{
		ChannelId: &channelId,
	}

	_, err := mediaLiveClient.StartChannel(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
