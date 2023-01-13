package commands

import (
	"github.com/nlm/briq-cli/config"
	"github.com/nlm/briq-cli/utils"
	"github.com/spf13/cobra"
)

const (
	viperKeyBriqSecretKey = "secret_key"
)

var (
	RootCmd = cobra.Command{
		Use:   "briq-cli",
		Short: "briq command-line utility",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(utils.PrefixError("config error", config.ConfigFile.Check()))
		},
	}
)

func Register(cmd *cobra.Command) {
	RootCmd.AddCommand(cmd)
}
