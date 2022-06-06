package main

import (
	"github.com/nlm/briq-cli/briq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	GiveBriqCmd.Flags().String("to", "", "username to give to")
	GiveBriqCmd.Flags().String("message", "Have a Briq! #Rock-solid", "message to send")
	GiveBriqCmd.Flags().Uint("count", 1, "how many briqs to give")
	GiveBriqCmd.MarkFlagRequired("to")
	Register(&GiveBriqCmd)
}

var GiveBriqCmd = cobra.Command{
	Use:     "give",
	Aliases: []string{"g"},
	Short:   "Give one or more briqs to someone",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		argTo, err := cmd.Flags().GetString("to")
		cobra.CheckErr(err)
		argMesg, err := cmd.Flags().GetString("message")
		cobra.CheckErr(err)
		client, err := briq.NewClient(viper.GetString(viperKeyBriqSecretKey))
		cobra.CheckErr(err)
		user, err := client.GetUser(
			cmd.Context(),
			&briq.GetUserRequest{Username: argTo},
		)
		cobra.CheckErr(err)
		argCount, err := cmd.Flags().GetUint("count")
		cobra.CheckErr(err)
		for i := uint(0); i < argCount && i < 10; i++ {
			req := &briq.CreateTransactionRequest{
				App:     briq.AppGive,
				Comment: argMesg,
				To:      user.Id,
			}
			res, err := client.CreateTransaction(cmd.Context(), req)
			cobra.CheckErr(err)
			cobra.CheckErr(Render(res))
		}
	},
}
