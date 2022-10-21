package harvester

import (
	"context"

	"github.com/harvester/terraform-provider-harvester/pkg/importer"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VirtualMachineGenerator struct {
	HarvesterService
}

func (g *VirtualMachineGenerator) InitResources() error {
	args := g.GetArgs()
	namespace := args[optionNamespace].(string)
	client, err := g.NewHarvesterClient()
	if err != nil {
		return err
	}

	virtualMachineList, err := client.HarvesterClient.KubevirtV1().VirtualMachines(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, vm := range virtualMachineList.Items {
		vmi, err := client.HarvesterClient.KubevirtV1().VirtualMachineInstances(vm.Namespace).Get(context.Background(), vm.Name, metav1.GetOptions{})
		if err != nil && !apierrors.IsNotFound(err) {
			return err
		}
		stateGetter, err := importer.ResourceVirtualMachineStateGetter(&vm, vmi, "")
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), stateGetter))
	}

	return nil
}
