package harvester

import (
	"context"

	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VLANConfigGenerator struct {
	HarvesterService
}

func (g *VLANConfigGenerator) InitResources() error {
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	vlanConfigList, err := client.HarvesterNetworkClient.NetworkV1beta1().VlanConfigs().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, vlanConfig := range vlanConfigList.Items {
		stateGetter, err := importer.ResourceVLANConfigStateGetter(&vlanConfig)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
