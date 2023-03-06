package harvester

import (
	"context"
	"fmt"

	"github.com/harvester/terraform-provider-harvester/pkg/client"
	"github.com/harvester/terraform-provider-harvester/pkg/constants"
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
		vmStateGetter, err := importer.ResourceVirtualMachineStateGetter(&vm, vmi, "")
		if err != nil {
			return err
		}

		g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), vmStateGetter)) // nolint
		secretStateGetters, err := getCloudInitSecretStateGetters(client, vm.Namespace, vmStateGetter)
		if err != nil {
			return err
		}
		for _, secretStateGetter := range secretStateGetters { // nolint
			g.Resources = append(g.Resources, buildResourceFormStateGetter(g.GetProviderName(), secretStateGetter)) // nolint
		}
	}

	return nil
}

func getCloudInitSecretStateGetters(client *client.Client, namespace string, vmStateGetter *importer.StateGetter) ([]*importer.StateGetter, error) {
	var val interface{}
	var ok bool
	if val, ok = vmStateGetter.States[constants.FieldVirtualMachineCloudInit]; !ok {
		return nil, nil
	}

	var cloudInit []map[string]interface{}
	if cloudInit, ok = val.([]map[string]interface{}); !ok {
		return nil, fmt.Errorf("cloudinit is not a list of map")
	}

	var userDataSecretName, networkDataSecretName string
	var secretStateGetters []*importer.StateGetter
	if val, ok = cloudInit[0][constants.FieldCloudInitUserDataSecretName]; ok {
		if userDataSecretName, ok = val.(string); !ok {
			return nil, fmt.Errorf("user_data_secret_name is not a string")
		}
		secret, err := client.KubeClient.CoreV1().Secrets(namespace).Get(context.Background(), userDataSecretName, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to get secret %s: %v", userDataSecretName, err)
		}
		secretStateGetter, err := importer.ResourceCloudInitSecretStateGetter(secret)
		if err != nil {
			return nil, fmt.Errorf("failed to get secret state getter: %v", err)
		}
		secretStateGetters = append(secretStateGetters, secretStateGetter)
	}
	if val, ok = cloudInit[0][constants.FieldCloudInitNetworkDataSecretName]; ok {
		if networkDataSecretName, ok = val.(string); !ok {
			return nil, fmt.Errorf("network_data_secret_name is not a string")
		}
		if networkDataSecretName == userDataSecretName {
			return secretStateGetters, nil
		}
		secret, err := client.KubeClient.CoreV1().Secrets(namespace).Get(context.Background(), networkDataSecretName, metav1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to get secret %s: %v", networkDataSecretName, err)
		}
		secretStateGetter, err := importer.ResourceCloudInitSecretStateGetter(secret)
		if err != nil {
			return nil, fmt.Errorf("failed to get secret state getter: %v", err)
		}
		secretStateGetters = append(secretStateGetters, secretStateGetter)
	}
	return secretStateGetters, nil
}
