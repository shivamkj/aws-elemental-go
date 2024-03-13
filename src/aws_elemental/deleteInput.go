package aws_elemental

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func DeleteInput(inputId string) {
	inputParam := medialive.DeleteInputInput{
		InputId: &inputId,
	}

	_, err := mediaLiveClient.DeleteInput(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
