package harvester

import (
	"context"

	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterNetworkGenerator struct {
	HarvesterService
}

func (g *ClusterNetworkGenerator) InitResources() error {
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	clusterNetworkList, err := client.HarvesterNetworkClient.NetworkV1beta1().ClusterNetworks().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, clusterNetwork := range clusterNetworkList.Items {
		stateGetter, err := importer.ResourceClusterNetworkStateGetter(&clusterNetwork)
		if err != nil {
			return nil
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
