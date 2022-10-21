package harvester

import (
	"context"

	"github.com/harvester/harvester/pkg/util"
	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StorageClassGenerator struct {
	HarvesterService
}

func (g *StorageClassGenerator) InitResources() error {
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	storageClassList, err := client.StorageClassClient.StorageClasses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, storageClass := range storageClassList.Items {
		if storageClass.Parameters[util.LonghornOptionBackingImageName] != "" {
			continue
		}

		stateGetter, err := importer.ResourceStorageClassStateGetter(&storageClass)
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
