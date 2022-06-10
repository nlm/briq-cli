package main

import (
	"fmt"

	"github.com/nlm/briq-cli/briq"
	"github.com/nlm/briq-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	RandomLoveCmd.Flags().String("to-group", "", "limit to a specific group of users")
	RandomLoveCmd.Flags().String("message", "Everybody needs some Briqs #Community", "message to send")
	RandomLoveCmd.Flags().Uint("user-count", 3, "number of users to send briqs to")
	RandomLoveCmd.Flags().Uint("amount", 1, "how many briqs to give")
	Register(&RandomLoveCmd)
}

const (
	randomLoveGroupName = "random-love"
)

var RandomLoveCmd = cobra.Command{
	Use:     "random-love",
	Aliases: []string{"rl"},
	Short:   "Send a briq to a number of random users",
	Example: "random-love --user-count 2 --amount 3 --to-group favorites",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		argToGroup, err := cmd.Flags().GetString("to-group")
		cobra.CheckErr(err)
		argMessage, err := cmd.Flags().GetString("message")
		cobra.CheckErr(err)
		argUserCount, err := cmd.Flags().GetUint("user-count")
		cobra.CheckErr(err)
		argAmount, err := cmd.Flags().GetUint("amount")
		cobra.CheckErr(err)
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		req := &briq.ListUsersRequest{}
		res, err := client.ListUsers(cmd.Context(), req)
		cobra.CheckErr(err)
		targetUsers := []briq.User(nil)
		if argToGroup == "" {
			// if a random-love group is defined, use it
			groupUsers := viper.GetStringSlice(fmt.Sprintf("groups.%s", randomLoveGroupName))
			if len(groupUsers) > 0 {
				targetUsers = utils.FilterSlice(res.Users, groupUsers, BriqUserKey)
			} else {
				targetUsers = res.Users
			}
		} else {
			// if a group is specified, only select users from this group
			groupUsers := viper.GetStringSlice(fmt.Sprintf("groups.%s", argToGroup))
			if len(groupUsers) == 0 {
				cobra.CheckErr(fmt.Errorf("group not found or empty"))
			}
			targetUsers = utils.FilterSlice(res.Users, groupUsers, BriqUserKey)
		}
		for _, user := range utils.RandSlice(targetUsers, int(argUserCount)) {
			for i := uint(0); i < CapBriqAmount(argAmount); i++ {
				fmt.Println("Sending gift to", user.Username)
				req := &briq.CreateTransactionRequest{
					App:     briq.AppGive,
					To:      user.Id,
					Comment: argMessage,
				}
				res, err := client.CreateTransaction(cmd.Context(), req)
				cobra.CheckErr(err)
				cobra.CheckErr(Render(res))
			}
		}
	},
}
