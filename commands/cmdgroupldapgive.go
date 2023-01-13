package commands

import (
	"fmt"

	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/render"
	"github.com/nlm/briq-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	GroupGiveLdapCmd.Flags().String("to-ldap-group", "", "group name")
	GroupGiveLdapCmd.Flags().String("message", "🎁 New feature working #Community", "message to send")
	GroupGiveLdapCmd.Flags().Uint("amount", 1, "how many briqs to give")
	GroupGiveLdapCmd.MarkFlagRequired("to-ldap-group")
	Register(&GroupGiveLdapCmd)
}

var GroupGiveLdapCmd = cobra.Command{
	Use:     "group-ldap-give",
	Aliases: []string{"glg"},
	Short:   "Give one or more briqs to a LDAP group",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		argToGroup, err := cmd.Flags().GetString("to-ldap-group")
		cobra.CheckErr(err)
		argMessage, err := cmd.Flags().GetString("message")
		cobra.CheckErr(err)
		argAmount, err := cmd.Flags().GetUint("amount")
		cobra.CheckErr(err)

		// Get members of group
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.GetGroupRequest{
			Name: argToGroup,
		}
		res, err := client.GetGroup(cmd.Context(), req)
		cobra.CheckErr(err)

		// Get current user username
		client, err = briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		reqMe := &briq.GetMeRequest{}
		resMe, err := client.GetMe(cmd.Context(), reqMe)
		cobra.CheckErr(err)

		for _, user := range res.Users {
			for i := uint(0); i < utils.CapBriqAmount(argAmount); i++ {
				if user.Username != resMe.Username {
					fmt.Printf("Sending gift to %s (%v)\n", user.Username, user.Id)
					req := &briq.CreateTransactionRequest{
						App:     briq.AppGive,
						To:      user.Id,
						Comment: argMessage,
					}
					res, err := client.CreateTransaction(cmd.Context(), req)
					cobra.CheckErr(err)
					cobra.CheckErr(render.Render(res))
				}
			}
		}
	},
}
