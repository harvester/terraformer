package harvester

import (
	"context"

	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	"github.com/rancher/wrangler/pkg/slice"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VolumeGenerator struct {
	HarvesterService
}

func (g *VolumeGenerator) InitResources() error {
	args := g.GetArgs()
	namespace := args[optionNamespace].(string)
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	volumeList, err := client.KubeClient.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, volume := range volumeList.Items {
		if slice.ContainsString([]string{"cattle-monitoring-system"}, volume.Namespace) {
			continue
		}

		stateGetter, err := importer.ResourceVolumeStateGetter(&volume)
		if err != nil {
			return nil
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
