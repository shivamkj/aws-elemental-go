package aws_elemental

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func StopChannel(channelId string) {
	inputParam := medialive.StopChannelInput{
		ChannelId: &channelId,
	}

	_, err := mediaLiveClient.StopChannel(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
