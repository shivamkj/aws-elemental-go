package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagev2/types"
)

var (
	segmentName  = "s-"
	menifestName = "LL-HLS"
)

func CreateOriginEndpoint(id string) string {
	var (
		channelName      = fmt.Sprintf("Channel-%s", id)
		channelGroupName = fmt.Sprintf("ChannelGroup-%s", id)
		endpointName     = fmt.Sprintf("Endpoint-%s", id)
	)
	inputParam := mediapackagev2.CreateOriginEndpointInput{
		ChannelName:        &channelName,
		ChannelGroupName:   &channelGroupName,
		OriginEndpointName: &endpointName,
		ContainerType:      types.ContainerTypeTs,
		Segment: &types.Segment{
			SegmentDurationSeconds: &segmentDuration,
			SegmentName:            &segmentName,
		},
		LowLatencyHlsManifests: []types.CreateLowLatencyHlsManifestConfiguration{
			{
				ManifestName: &menifestName,
			},
		},
		Tags: commonTag,
	}

	output, err := mediaPackageClient.CreateOriginEndpoint(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}

	return *output.LowLatencyHlsManifests[0].Url
}
