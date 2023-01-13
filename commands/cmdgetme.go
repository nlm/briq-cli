package commands

import (
	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/render"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Register(&GetMeCmd)
}

var GetMeCmd = cobra.Command{
	Use:     "me",
	Aliases: []string{"m"},
	Short:   "Get information about me",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.GetMeRequest{}
		res, err := client.GetMe(cmd.Context(), req)
		cobra.CheckErr(err)
		cobra.CheckErr(render.Render(res))
	},
}
