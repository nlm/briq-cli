package commands

import (
	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/render"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	GetGroupCmd.Flags().String("group", "", "group name")
	GetGroupCmd.MarkFlagRequired("group")
	Register(&GetGroupCmd)
}

var GetGroupCmd = cobra.Command{
	Use:     "group",
	Aliases: []string{"g"},
	Short:   "Get info about a LDAP group",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		argGroup, err := cmd.Flags().GetString("group")
		cobra.CheckErr(err)
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.GetGroupRequest{
			Name: argGroup,
		}
		res, err := client.GetGroup(cmd.Context(), req)
		cobra.CheckErr(err)
		cobra.CheckErr(render.Render(res))
	},
}
