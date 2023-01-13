package commands

import (
	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/render"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Register(&GetGroupsCmd)
}

var GetGroupsCmd = cobra.Command{
	Use:     "list-ldap-groups",
	Aliases: []string{"lg"},
	Short:   "List LDAP groups",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.ListGroupsRequest{}
		res, err := client.ListGroups(cmd.Context(), req)
		cobra.CheckErr(err)
		cobra.CheckErr(render.Render(res))
	},
}
