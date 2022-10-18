package harvester

import (
	"errors"
	"os"

	"github.com/zclconf/go-cty/cty"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

const (
	optionKubeConfig  = "kubeconfig"
	optionKubeContext = "kubecontext"
	optionNamespace   = "namespace"

	envKubeConfig  = "KUBECONFIG"
	envKubeContext = "KUBECONTEXT"
	envNamespace   = "NAMESPACE"

	ProviderName = "harvester"
)

type HarvesterProvider struct { //nolint
	terraformutils.Provider
	kubeconfig  string
	kubecontext string
	namespace   string
}

func (p *HarvesterProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *HarvesterProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			p.GetName(): map[string]interface{}{
				optionKubeConfig:  p.kubeconfig,
				optionKubeContext: p.kubecontext,
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
	p.kubecontext = os.Getenv(envKubeContext)
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
		optionKubeConfig:  p.kubeconfig,
		optionKubeContext: p.kubecontext,
		optionNamespace:   p.namespace,
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
		"storageclass":   &StorageClassGenerator{},
		"vlanconfig":     &VLANConfigGenerator{},
	}
}
