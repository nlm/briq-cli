package main

import (
	"log"

	"github.com/nlm/briq-cli/briq"
	"github.com/spf13/cobra"
)

func init() {
	Register(&RandomLoveCmd)
}

var RandomLoveCmd = cobra.Command{
	Use:     "random-love",
	Aliases: []string{"rl"},
	Run: func(cmd *cobra.Command, args []string) {
		client, err := briq.NewClient(briqSecretKey)
		if err != nil {
			log.Fatal(err)
		}
		req := &briq.GetUsersRequest{}
		res, err := client.GetUsers(cmd.Context(), req)
		if err != nil {
			log.Fatal(err)
		}
		randUsers, err := RandSlice(res.Users, 5)
		if err != nil {
			log.Fatal(err)
		}
		for _, user := range randUsers {
			log.Println("Sending gift to", user.DisplayName)
			req := &briq.CreateTransactionRequest{
				App:     briq.AppGive,
				To:      user.Id,
				Comment: "Random Gift #Community",
			}
			res, err := client.CreateTransaction(cmd.Context(), req)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
	},
}
