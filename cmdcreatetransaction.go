package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/nlm/briq-cli/briq"
	"github.com/spf13/cobra"
)

func init() {
	SendBriqCmd.Flags().String("to", "", "uuid to give to")
	SendBriqCmd.Flags().String("message", "Have a Briq! #Rock-solid", "message to send")
	Register(&SendBriqCmd)
}

var SendBriqCmd = cobra.Command{
	Use:     "send-briq",
	Aliases: []string{"send"},
	Run: func(cmd *cobra.Command, args []string) {
		argTo, _ := cmd.Flags().GetString("to")
		argMesg, _ := cmd.Flags().GetString("message")
		uid, err := uuid.Parse(argTo)
		if err != nil {
			cmd.Usage()
			log.Fatal(err)
		}
		client, err := briq.NewClient(briqSecretKey)
		if err != nil {
			log.Fatal(err)
		}
		req := &briq.CreateTransactionRequest{
			App:     briq.AppGive,
			Comment: argMesg,
			To:      uid,
		}
		res, err := client.CreateTransaction(cmd.Context(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)
	},
}
