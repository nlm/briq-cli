package main

import (
	"fmt"
	"log"

	"github.com/nlm/briq-cli/briq"
	"github.com/spf13/cobra"
)

func init() {
	Register(&GetUsersCmd)
}

var GetUsersCmd = cobra.Command{
	Use:     "get-users",
	Aliases: []string{"gu"},
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
		for _, user := range res.Users {
			fmt.Println(user)
		}
	},
}
