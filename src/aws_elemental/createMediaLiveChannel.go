package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
)

var (
	inputName     = "RTMP-Input"
	destinationId = "destination1"
)

func CreateMediaLiveChannel(id, inputId string, ingestEndpoints []string) string {
	destinationName := fmt.Sprintf("MediaLive Channel %s", id)
	inputParam := medialive.CreateChannelInput{
		ChannelClass: types.ChannelClassStandard,
		Destinations: []types.OutputDestination{
			{
				Id: &destinationId,
				Settings: []types.OutputDestinationSettings{
					{
						Url: &ingestEndpoints[0],
					},
					{
						Url: &ingestEndpoints[1],
					},
				},
			},
		},
		InputAttachments: []types.InputAttachment{
			{
				InputAttachmentName: &inputName,
				InputId:             &inputId,
			},
		},
		InputSpecification: &types.InputSpecification{
			Codec:          types.InputCodecAvc,
			MaximumBitrate: types.InputMaximumBitrateMax10Mbps,
			Resolution:     types.InputResolutionSd,
		},
		EncoderSettings: getEncoderSettings(),
		LogLevel:        types.LogLevelError,
		// Maintenance: &types.MaintenanceCreateSettings{
		// 	MaintenanceDay: types.MaintenanceDaySaturday,
		// 	MaintenanceStartTime: "",
		// },
		Name:    &destinationName,
		RoleArn: &roleArn,
		Tags:    commonTag,
	}

	output, err := mediaLiveClient.CreateChannel(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}

	return *output.Channel.Id
}

func getEncoderSettings() *types.EncoderSettings {
	var (
		audioName, audioSelectorName       = "audio_1_aac64", "default"
		videoName, _                       = "video_416_234", ""
		destinationRef                     = "destination1"
		videoHeight, videoWidth      int32 = 236, 416
	)

	encoderSettings := types.EncoderSettings{
		TimecodeConfig: &types.TimecodeConfig{
			Source: types.TimecodeConfigSourceSystemclock,
		},
		AudioDescriptions: []types.AudioDescription{
			{
				Name:              &audioName,
				AudioSelectorName: &audioSelectorName,
			},
		},
		OutputGroups: []types.OutputGroup{
			{
				OutputGroupSettings: &types.OutputGroupSettings{
					HlsGroupSettings: &types.HlsGroupSettings{
						Destination: &types.OutputLocationRef{
							DestinationRefId: &destinationRef,
						},
						HlsCdnSettings: &types.HlsCdnSettings{
							HlsBasicPutSettings: &types.HlsBasicPutSettings{},
						},
						SegmentLength: &segmentDuration,
					},
				},
				Outputs: []types.Output{
					{
						AudioDescriptionNames: []string{
							audioName,
						},
						VideoDescriptionName: &videoName,
						OutputSettings: &types.OutputSettings{
							HlsOutputSettings: &types.HlsOutputSettings{
								HlsSettings: &types.HlsSettings{
									StandardHlsSettings: &types.StandardHlsSettings{
										M3u8Settings: &types.M3u8Settings{},
									},
								},
							},
						},
					},
				},
			},
		},
		VideoDescriptions: []types.VideoDescription{
			{
				Name: &videoName,
				CodecSettings: &types.VideoCodecSettings{
					H264Settings: &types.H264Settings{},
				},
				Height: &videoHeight,
				Width:  &videoWidth,
			},
		},
	}

	return &encoderSettings
}

// const ENCODER_SETTINGS = {
//     "AudioDescriptions": [
//         {
//             "AudioSelectorName": "default",
//             "AudioTypeControl": "FOLLOW_INPUT",
//             "CodecSettings": {
//                 "AacSettings": {
//                     "Bitrate": 64000,
//                     "RawFormat": "NONE",
//                     "Spec": "MPEG4"
//                 }
//             },
//             "LanguageCodeControl": "FOLLOW_INPUT",
//             "Name": "audio_1_aac64"
//         },
//         {
//             "AudioSelectorName": "default",
//             "AudioTypeControl": "FOLLOW_INPUT",
//             "CodecSettings": {
//                 "AacSettings": {
//                     "Bitrate": 64000,
//                     "RawFormat": "NONE",
//                     "Spec": "MPEG4"
//                 }
//             },
//             "LanguageCodeControl": "FOLLOW_INPUT",
//             "Name": "audio_3_aac64"
//         },
//         {
//             "AudioSelectorName": "default",
//             "AudioTypeControl": "FOLLOW_INPUT",
//             "CodecSettings": {
//                 "AacSettings": {
//                     "Bitrate": 96000,
//                     "RawFormat": "NONE",
//                     "Spec": "MPEG4"
//                 }
//             },
//             "LanguageCodeControl": "FOLLOW_INPUT",
//             "Name": "audio_2_aac96"
//         },
//         {
//             "AudioSelectorName": "default",
//             "AudioTypeControl": "FOLLOW_INPUT",
//             "CodecSettings": {
//                 "AacSettings": {
//                     "Bitrate": 96000,
//                     "RawFormat": "NONE",
//                     "Spec": "MPEG4"
//                 }
//             },
//             "LanguageCodeControl": "FOLLOW_INPUT",
//             "Name": "audio_3_aac96"
//         }
//     ],
//     "CaptionDescriptions": [],
//     "OutputGroups": [
//         {
//             "Name": "TN2224",
//             "OutputGroupSettings": {
//                 "HlsGroupSettings": {
//                     "AdMarkers": [],
//                     "CaptionLanguageMappings": [],
//                     "CaptionLanguageSetting": "OMIT",
//                     "ClientCache": "ENABLED",
//                     "CodecSpecification": "RFC_4281",
//                     "Destination": {
//                         "DestinationRefId": "destination1"
//                     },
//                     "DirectoryStructure": "SINGLE_DIRECTORY",
//                     "HlsCdnSettings": {
//                         "HlsBasicPutSettings": {
//                             "ConnectionRetryInterval": 30,
//                             "FilecacheDuration": 300,
//                             "NumRetries": 5,
//                             "RestartDelay": 5
//                         }
//                     },
//                     "IndexNSegments": 10,
//                     "InputLossAction": "EMIT_OUTPUT",
//                     "IvInManifest": "INCLUDE",
//                     "IvSource": "FOLLOWS_SEGMENT_NUMBER",
//                     "KeepSegments": 21,
//                     "ManifestCompression": "NONE",
//                     "ManifestDurationFormat": "FLOATING_POINT",
//                     "Mode": "LIVE",
//                     "OutputSelection": "MANIFESTS_AND_SEGMENTS",
//                     "ProgramDateTime": "INCLUDE",
//                     "ProgramDateTimePeriod": 600,
//                     "SegmentLength": 4,
//                     "SegmentationMode": "USE_SEGMENT_DURATION",
//                     "SegmentsPerSubdirectory": 10000,
//                     "StreamInfResolution": "INCLUDE",
//                     "TimedMetadataId3Frame": "PRIV",
//                     "TimedMetadataId3Period": 10,
//                     "TsFileMode": "SEGMENTED_FILES"
//                 }
//             },
//             "Outputs": [
//                 {
//                     "AudioDescriptionNames": [
//                         "audio_2_aac96"
//                     ],
//                     "CaptionDescriptionNames": [],
//                     "OutputSettings": {
//                         "HlsOutputSettings": {
//                             "H265PackagingType": "HVC1",
//                             "HlsSettings": {
//                                 "StandardHlsSettings": {
//                                     "AudioRenditionSets": "program_audio",
//                                     "M3u8Settings": {
//                                         "AudioFramesPerPes": 4,
//                                         "AudioPids": "492-498",
//                                         "EcmPid": "8182",
//                                         "KlvBehavior": "NO_PASSTHROUGH",
//                                         "NielsenId3Behavior": "NO_PASSTHROUGH",
//                                         "PcrControl": "PCR_EVERY_PES_PACKET",
//                                         "PmtPid": "480",
//                                         "ProgramNum": 1,
//                                         "Scte35Behavior": "NO_PASSTHROUGH",
//                                         "Scte35Pid": "500",
//                                         "TimedMetadataBehavior": "NO_PASSTHROUGH",
//                                         "TimedMetadataPid": "502",
//                                         "VideoPid": "481"
//                                     }
//                                 }
//                             },
//                             "NameModifier": "_960x540_2000k"
//                         }
//                     },
//                     "VideoDescriptionName": "video_960_540"
//                 },
//                 {
//                     "AudioDescriptionNames": [
//                         "audio_3_aac96"
//                     ],
//                     "CaptionDescriptionNames": [],
//                     "OutputSettings": {
//                         "HlsOutputSettings": {
//                             "HlsSettings": {
//                                 "StandardHlsSettings": {
//                                     "AudioRenditionSets": "program_audio",
//                                     "M3u8Settings": {
//                                         "AudioPids": "492-498",
//                                         "EcmPid": "8182",
//                                         "PcrControl": "PCR_EVERY_PES_PACKET",
//                                         "PmtPid": "480",
//                                         "Scte35Behavior": "NO_PASSTHROUGH",
//                                         "Scte35Pid": "500",
//                                         "TimedMetadataBehavior": "NO_PASSTHROUGH",
//                                         "VideoPid": "481"
//                                     }
//                                 }
//                             },
//                             "NameModifier": "_1280x720_3300k"
//                         }
//                     },
//                     "VideoDescriptionName": "video_1280_720_1"
//                 },
//                 {
//                     "AudioDescriptionNames": [
//                         "audio_1_aac64"
//                     ],
//                     "CaptionDescriptionNames": [],
//                     "OutputSettings": {
//                         "HlsOutputSettings": {
//                             "HlsSettings": {
//                                 "StandardHlsSettings": {
//                                     "AudioRenditionSets": "program_audio",
//                                     "M3u8Settings": {
//                                         "AudioPids": "492-498",
//                                         "EcmPid": "8182",
//                                         "PcrControl": "PCR_EVERY_PES_PACKET",
//                                         "PmtPid": "480",
//                                         "Scte35Behavior": "NO_PASSTHROUGH",
//                                         "Scte35Pid": "500",
//                                         "TimedMetadataBehavior": "NO_PASSTHROUGH",
//                                         "VideoPid": "481"
//                                     }
//                                 }
//                             },
//                             "NameModifier": "_416x234_200k"
//                         }
//                     },
//                     "VideoDescriptionName": "video_416_234"
//                 },
//                 {
//                     "AudioDescriptionNames": [
//                         "audio_3_aac64"
//                     ],
//                     "CaptionDescriptionNames": [],
//                     "OutputSettings": {
//                         "HlsOutputSettings": {
//                             "H265PackagingType": "HVC1",
//                             "HlsSettings": {
//                                 "StandardHlsSettings": {
//                                     "AudioRenditionSets": "program_audio",
//                                     "M3u8Settings": {
//                                         "AudioFramesPerPes": 4,
//                                         "AudioPids": "492-498",
//                                         "EcmPid": "8182",
//                                         "KlvBehavior": "NO_PASSTHROUGH",
//                                         "NielsenId3Behavior": "NO_PASSTHROUGH",
//                                         "PcrControl": "PCR_EVERY_PES_PACKET",
//                                         "PmtPid": "480",
//                                         "ProgramNum": 1,
//                                         "Scte35Behavior": "NO_PASSTHROUGH",
//                                         "Scte35Pid": "500",
//                                         "TimedMetadataBehavior": "NO_PASSTHROUGH",
//                                         "TimedMetadataPid": "502",
//                                         "VideoPid": "481"
//                                     }
//                                 }
//                             },
//                             "NameModifier": "_640x360_800k"
//                         }
//                     },
//                     "VideoDescriptionName": "video_640_360"
//                 }
//             ]
//         }
//     ],
//     "TimecodeConfig": {
//         "Source": "SYSTEMCLOCK"
//     },
//     "VideoDescriptions": [
//         {
//             "CodecSettings": {
//                 "H264Settings": {
//                     "AdaptiveQuantization": "HIGH",
//                     "Bitrate": 200000,
//                     "ColorMetadata": "INSERT",
//                     "EntropyEncoding": "CAVLC",
//                     "FlickerAq": "ENABLED",
//                     "FramerateControl": "SPECIFIED",
//                     "FramerateDenominator": 1001,
//                     "FramerateNumerator": 15000,
//                     "GopBReference": "DISABLED",
//                     "GopNumBFrames": 0,
//                     "GopSize": 30,
//                     "GopSizeUnits": "FRAMES",
//                     "Level": "H264_LEVEL_3",
//                     "LookAheadRateControl": "HIGH",
//                     "ParControl": "INITIALIZE_FROM_SOURCE",
//                     "Profile": "BASELINE",
//                     "RateControlMode": "CBR",
//                     "SceneChangeDetect": "ENABLED",
//                     "SpatialAq": "ENABLED",
//                     "Syntax": "DEFAULT",
//                     "TemporalAq": "ENABLED"
//                 }
//             },
//             "Height": 236,
//             "Name": "video_416_234",
//             "ScalingBehavior": "DEFAULT",
//             "Width": 416
//         },
//         {
//             "CodecSettings": {
//                 "H264Settings": {
//                     "AdaptiveQuantization": "HIGH",
//                     "AfdSignaling": "NONE",
//                     "Bitrate": 800000,
//                     "ColorMetadata": "INSERT",
//                     "EntropyEncoding": "CABAC",
//                     "FlickerAq": "ENABLED",
//                     "ForceFieldPictures": "DISABLED",
//                     "FramerateControl": "SPECIFIED",
//                     "FramerateDenominator": 1001,
//                     "FramerateNumerator": 30000,
//                     "GopBReference": "ENABLED",
//                     "GopClosedCadence": 1,
//                     "GopNumBFrames": 3,
//                     "GopSize": 60,
//                     "GopSizeUnits": "FRAMES",
//                     "Level": "H264_LEVEL_3",
//                     "LookAheadRateControl": "HIGH",
//                     "NumRefFrames": 1,
//                     "ParControl": "INITIALIZE_FROM_SOURCE",
//                     "Profile": "MAIN",
//                     "RateControlMode": "CBR",
//                     "ScanType": "PROGRESSIVE",
//                     "SceneChangeDetect": "ENABLED",
//                     "SpatialAq": "ENABLED",
//                     "SubgopLength": "FIXED",
//                     "Syntax": "DEFAULT",
//                     "TemporalAq": "ENABLED",
//                     "TimecodeInsertion": "DISABLED"
//                 }
//             },
//             "Height": 360,
//             "Name": "video_640_360",
//             "RespondToAfd": "NONE",
//             "ScalingBehavior": "DEFAULT",
//             "Sharpness": 50,
//             "Width": 640
//         },
//         {
//             "CodecSettings": {
//                 "H264Settings": {
//                     "AdaptiveQuantization": "HIGH",
//                     "AfdSignaling": "NONE",
//                     "Bitrate": 2200000,
//                     "ColorMetadata": "INSERT",
//                     "EntropyEncoding": "CABAC",
//                     "FlickerAq": "ENABLED",
//                     "ForceFieldPictures": "DISABLED",
//                     "FramerateControl": "SPECIFIED",
//                     "FramerateDenominator": 1001,
//                     "FramerateNumerator": 30000,
//                     "GopBReference": "ENABLED",
//                     "GopClosedCadence": 1,
//                     "GopNumBFrames": 3,
//                     "GopSize": 60,
//                     "GopSizeUnits": "FRAMES",
//                     "Level": "H264_LEVEL_4_1",
//                     "LookAheadRateControl": "HIGH",
//                     "NumRefFrames": 1,
//                     "ParControl": "INITIALIZE_FROM_SOURCE",
//                     "Profile": "HIGH",
//                     "RateControlMode": "CBR",
//                     "ScanType": "PROGRESSIVE",
//                     "SceneChangeDetect": "ENABLED",
//                     "SpatialAq": "ENABLED",
//                     "SubgopLength": "FIXED",
//                     "Syntax": "DEFAULT",
//                     "TemporalAq": "ENABLED",
//                     "TimecodeInsertion": "DISABLED"
//                 }
//             },
//             "Height": 540,
//             "Name": "video_960_540",
//             "RespondToAfd": "NONE",
//             "ScalingBehavior": "DEFAULT",
//             "Sharpness": 50,
//             "Width": 960
//         },
//         {
//             "CodecSettings": {
//                 "H264Settings": {
//                     "AdaptiveQuantization": "HIGH",
//                     "Bitrate": 3300000,
//                     "ColorMetadata": "INSERT",
//                     "EntropyEncoding": "CABAC",
//                     "FlickerAq": "ENABLED",
//                     "FramerateControl": "SPECIFIED",
//                     "FramerateDenominator": 1001,
//                     "FramerateNumerator": 30000,
//                     "GopBReference": "ENABLED",
//                     "GopNumBFrames": 3,
//                     "GopSize": 60,
//                     "GopSizeUnits": "FRAMES",
//                     "Level": "H264_LEVEL_4_1",
//                     "LookAheadRateControl": "HIGH",
//                     "ParControl": "INITIALIZE_FROM_SOURCE",
//                     "Profile": "HIGH",
//                     "RateControlMode": "CBR",
//                     "SceneChangeDetect": "ENABLED",
//                     "SpatialAq": "ENABLED",
//                     "Syntax": "DEFAULT",
//                     "TemporalAq": "ENABLED"
//                 }
//             },
//             "Height": 720,
//             "Name": "video_1280_720_1",
//             "ScalingBehavior": "DEFAULT",
//             "Width": 1280
//         }
//     ]
// }
