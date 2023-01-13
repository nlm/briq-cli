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
	GroupGiveCmd.Flags().String("to-group", "", "group name")
	GroupGiveCmd.Flags().String("message", "🎁 #Community", "message to send")
	GroupGiveCmd.Flags().Uint("amount", 1, "how many briqs to give")
	GroupGiveCmd.MarkFlagRequired("group")
	Register(&GroupGiveCmd)
}

var GroupGiveCmd = cobra.Command{
	Use:     "group-give",
	Aliases: []string{"gg"},
	Short:   "Give one or more briqs to a group",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		argToGroup, err := cmd.Flags().GetString("to-group")
		cobra.CheckErr(err)
		argMessage, err := cmd.Flags().GetString("message")
		cobra.CheckErr(err)
		argAmount, err := cmd.Flags().GetUint("amount")
		cobra.CheckErr(err)
		groupUserNames := viper.GetStringSlice(fmt.Sprintf("groups.%s", argToGroup))
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.ListUsersRequest{}
		res, err := client.ListUsers(cmd.Context(), req)
		cobra.CheckErr(err)
		for _, user := range utils.FilterSlice(res.Users, groupUserNames, utils.BriqUserKey) {
			for i := uint(0); i < utils.CapBriqAmount(argAmount); i++ {
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
	},
}
