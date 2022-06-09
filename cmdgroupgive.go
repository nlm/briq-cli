package main

import (
	"fmt"

	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	MyRoundCmd.Flags().String("group", "", "group name")
	MyRoundCmd.Flags().String("message", "🎁 #Community", "message to send")
	MyRoundCmd.MarkFlagRequired("group")
	Register(&MyRoundCmd)
}

var MyRoundCmd = cobra.Command{
	Use:     "group-give",
	Aliases: []string{"gg"},
	Short:   "Give one or more briqs to a group",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		groupName, err := cmd.Flags().GetString("group")
		cobra.CheckErr(err)
		message, err := cmd.Flags().GetString("message")
		cobra.CheckErr(err)
		groupUserNames := viper.GetStringSlice(fmt.Sprintf("groups.%s", groupName))
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.ListUsersRequest{}
		res, err := client.ListUsers(cmd.Context(), req)
		cobra.CheckErr(err)
		for _, user := range utils.FilterSlice(res.Users, groupUserNames, BriqUserKey) {
			fmt.Printf("Sending gift to %s (%v)\n", user.Username, user.Id)
			req := &briq.CreateTransactionRequest{
				App:     briq.AppGive,
				To:      user.Id,
				Comment: message,
			}
			res, err := client.CreateTransaction(cmd.Context(), req)
			cobra.CheckErr(err)
			cobra.CheckErr(Render(res))
		}
	},
}
