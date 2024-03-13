package aws_elemental

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/shivamkj/live-streaming-service/src/utils"
)

func CreateInputSecurityGroup(id string, cidr string) string {
	var tags = utils.DeepCopyMap(commonTag)
	tags["id"] = id

	inputParam := medialive.CreateInputSecurityGroupInput{
		WhitelistRules: []types.InputWhitelistRuleCidr{
			{Cidr: &cidr},
		},
		Tags: tags,
	}

	output, err := mediaLiveClient.CreateInputSecurityGroup(context.TODO(), &inputParam)
	if err != nil {
		panic(err)
	}

	return *output.SecurityGroup.Id
}
