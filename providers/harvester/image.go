package harvester

import (
	"context"

	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ImageGenerator struct {
	HarvesterService
}

func (g *ImageGenerator) InitResources() error {
	args := g.GetArgs()
	namespace := args[optionNamespace].(string)
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	imageList, err := client.HarvesterClient.HarvesterhciV1beta1().VirtualMachineImages(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, image := range imageList.Items {
		stateGetter, err := importer.ResourceImageStateGetter(&image)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
