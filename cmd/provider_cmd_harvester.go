package cmd

import (
	"github.com/spf13/cobra"

	"github.com/GoogleCloudPlatform/terraformer/providers/harvester"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

func newCmdHarvesterImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "harvester",
		Short: "Import current state to Terraform configuration from Harvester",
		Long:  "Import current state to Terraform configuration from Harvester",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newHarvesterProvider()
			err := Import(provider, options, []string{})
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.AddCommand(listCmd(newHarvesterProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "virtualmachine", "virtualmachine=id1:id2:id3, id format: namespace/name")

	return cmd
}

func newHarvesterProvider() terraformutils.ProviderGenerator {
	return &harvester.HarvesterProvider{}
}
