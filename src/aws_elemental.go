package live_streaming

import (
	"time"

	elemental "github.com/shivamkj/live-streaming-service/src/aws_elemental"
	. "github.com/shivamkj/live-streaming-service/src/utils"
)

type AwsElemental struct {
	Id string
}

func (e AwsElemental) CreateStream() {
	const id = "9nibxugb"

	// Create a Media Live Input
	securityGroupId := elemental.CreateInputSecurityGroup(id, "0.0.0.0/0")
	input := elemental.CreateInput(id, securityGroupId)
	Logger.Debug("media live input created", "inputId", input.Id, "Url1", input.Url1, "Url2", input.Url2, "securityGroupId", securityGroupId)

	// Create Media Package for Output
	_ = elemental.CreateMediaPackageChannelGroup(id)
	ingestEndpoints := elemental.CreateMediaPackageChannel(id)
	videoUrl := elemental.CreateOriginEndpoint(id)
	Logger.Debug("media package channel group, channel & endpoint created", "ingestEndpoints", ingestEndpoints, "videoUrl", videoUrl)

	elemental.PutOriginEndpointPolicy(id) // Make Media Package Output Public
	Logger.Debug("origin endpoint policy created")

	channelId := elemental.CreateMediaLiveChannel(id, input.Id, ingestEndpoints)
	Logger.Debug("media live channel created", "channelId", channelId)
}

func (e AwsElemental) StartStream() {
	elemental.StartChannel("501146")
	Logger.Debug("channel started", "channelId", "501146")
}

func (e AwsElemental) StopStream() {
	const id = "9nibxugb"

	elemental.StopChannel("501146")
	time.Sleep(10 * time.Second)
	elemental.DeleteMediaLiveChannel("501146")

	time.Sleep(30 * time.Second)
	elemental.DeleteInput("2568422")
	time.Sleep(10 * time.Second)
	elemental.DeleteInputSecurityGroup("6881973")
	Logger.Debug("media live input & security group deleted")

	elemental.DeleteOriginEndpoint(id)
	elemental.DeleteMediaPackageChannel(id)
	elemental.DeleteMediaPackageChannelGroup(id)
	Logger.Debug("media package channel group, channel & endpoint deleted")
}

func (e AwsElemental) ListStream() error {
	return nil
}

func (e AwsElemental) GetStreamDetails() error {
	return nil
}
