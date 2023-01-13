package commands

import (
	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/render"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	Register(&ListUsersCmd)
}

var ListUsersCmd = cobra.Command{
	Use:     "list-users",
	Aliases: []string{"lu"},
	Short:   "List existing users on briq",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.ListUsersRequest{}
		res, err := client.ListUsers(cmd.Context(), req)
		cobra.CheckErr(err)
		cobra.CheckErr(render.Render(res))
	},
}
