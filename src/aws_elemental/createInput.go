package aws_elemental

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
)

type InputResult struct {
	Id   string
	Url1 string
	Url2 string
}

func CreateInput(id string, securityGroupId string) InputResult {
	var (
		inputName  = fmt.Sprintf("input-%s", id)
		streamName = fmt.Sprintf("live/%s", id)
	)
	inputParam := medialive.CreateInputInput{
		Destinations: []types.InputDestinationRequest{
			{StreamName: &streamName},
			{StreamName: &streamName},
		},
		InputSecurityGroups: []string{
			securityGroupId,
		},
		Name: &inputName,
		Type: types.InputTypeRtmpPush,
		Tags: commonTag,
	}

	output, err := mediaLiveClient.CreateInput(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}

	return InputResult{
		*output.Input.Id,
		*output.Input.Destinations[0].Url,
		*output.Input.Destinations[1].Url,
	}
}
