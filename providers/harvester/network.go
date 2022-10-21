package harvester

import (
	"context"

	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkGenerator struct {
	HarvesterService
}

func (g *NetworkGenerator) InitResources() error {
	args := g.GetArgs()
	namespace := args[optionNamespace].(string)
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	networkList, err := client.HarvesterClient.K8sCniCncfIoV1().NetworkAttachmentDefinitions(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, network := range networkList.Items {
		stateGetter, err := importer.ResourceNetworkStateGetter(&network)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
