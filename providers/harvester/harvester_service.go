package harvester

import (
	"fmt"

	"github.com/harvester/terraform-provider-harvester/pkg/client"
	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	"github.com/hashicorp/terraform/terraform"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type HarvesterService struct { //nolint
	terraformutils.Service
}

func (s *HarvesterService) NewHarvesterClient() (*client.Client, error) {
	args := s.GetArgs()
	return client.NewClient(args[optionKubeConfig].(string), args[optionKubeContext].(string))
}

func buildResourceFormStateGetter(providerName string, getter *importer.StateGetter) terraformutils.Resource {
	return terraformutils.Resource{
		ResourceName: getter.Name,
		Item:         nil,
		Provider:     providerName,
		InstanceState: &terraform.InstanceState{
			ID: getter.ID,
		},
		InstanceInfo: &terraform.InstanceInfo{
			Type: getter.ResourceType,
			Id:   fmt.Sprintf("%s.%s", getter.ResourceType, getter.Name),
		},
	}
}
