package harvester

import (
	"context"

	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KeyPairGenerator struct {
	HarvesterService
}

func (g *KeyPairGenerator) InitResources() error {
	args := g.GetArgs()
	namespace := args[optionNamespace].(string)
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	keypairList, err := client.HarvesterClient.HarvesterhciV1beta1().KeyPairs(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, keypair := range keypairList.Items {
		stateGetter, err := importer.ResourceKeyPairStateGetter(&keypair)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
