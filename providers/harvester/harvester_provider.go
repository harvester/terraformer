package harvester

import (
	"errors"
	"os"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

const (
	optionKubeConfig = "kubeconfig"
	optionNamespace  = "namespace"

	envKubeConfig = "KUBECONFIG"
	envNamespace  = "NAMESPACE"

	ProviderName = "harvester"
)

type HarvesterProvider struct { //nolint
	terraformutils.Provider
	kubeconfig string
	namespace  string
}

func (p *HarvesterProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *HarvesterProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				optionKubeConfig: p.kubeconfig,
			},
		},
	}
}

func (p *HarvesterProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		optionKubeConfig: cty.StringVal(p.kubeconfig),
	})
}

func (p *HarvesterProvider) Init(args []string) error {
	p.kubeconfig = os.Getenv(envKubeConfig)
	p.namespace = os.Getenv(envNamespace)
	return nil
}

func (p *HarvesterProvider) GetName() string {
	return ProviderName
}

func (p *HarvesterProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(p.GetName() + ": " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		optionKubeConfig: p.kubeconfig,
		optionNamespace:  p.namespace,
	})
	return nil
}

func (p *HarvesterProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"image":          &ImageGenerator{},
		"network":        &NetworkGenerator{},
		"volume":         &VolumeGenerator{},
		"clusternetwork": &ClusterNetworkGenerator{},
		"ssh_key":        &KeyPairGenerator{},
		"virtualmachine": &VirtualMachineGenerator{},
	}
}
