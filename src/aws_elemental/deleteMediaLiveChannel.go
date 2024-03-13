package aws_elemental

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func DeleteMediaLiveChannel(channelId string) {
	inputParam := medialive.DeleteChannelInput{
		ChannelId: &channelId,
	}

	_, err := mediaLiveClient.DeleteChannel(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
