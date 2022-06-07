package main

import (
	"fmt"

	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RandomLoveCmd.Flags().Uint("count", 5, "number of users to send briqs to")
	RandomLoveCmd.Flags().String("message", "Everybody needs some Briqs #Community", "message to send")
	Register(&RandomLoveCmd)
}

var RandomLoveCmd = cobra.Command{
	Use:     "random-love",
	Aliases: []string{"rl"},
	Short:   "Send a briq to a number of random users",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		count, err := cmd.Flags().GetUint("count")
		cobra.CheckErr(err)
		message, err := cmd.Flags().GetString("message")
		cobra.CheckErr(err)
		req := &briq.ListUsersRequest{}
		res, err := client.ListUsers(cmd.Context(), req)
		cobra.CheckErr(err)
		targetUsers := []briq.User(nil)
		lovedUsers := viper.GetStringSlice(viperKeyBriqLovedUsers)
		if len(lovedUsers) == 0 {
			targetUsers = res.Users
		} else {
			targetUsers = utils.FilterSlice(res.Users, lovedUsers, func(user briq.User) string { return user.Username })
		}
		for _, user := range utils.RandSlice(targetUsers, int(count)) {
			fmt.Println("Sending gift to", user.Username)
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
