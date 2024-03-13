package aws_elemental

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
)

func DeleteInputSecurityGroup(securityGroupId string) {
	inputParam := medialive.DeleteInputSecurityGroupInput{
		InputSecurityGroupId: &securityGroupId,
	}

	_, err := mediaLiveClient.DeleteInputSecurityGroup(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}
}
