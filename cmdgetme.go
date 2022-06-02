package main

import (
	"log"

	"github.com/nlm/briq-cli/briq"
	"github.com/spf13/cobra"
)

func init() {
	Register(&GetMeCmd)
}

var GetMeCmd = cobra.Command{
	Use: "me",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := briq.NewClient(briqSecretKey)
		if err != nil {
			log.Fatal(err)
		}
		req := &briq.GetMeRequest{}
		res, err := client.GetMe(cmd.Context(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)
	},
}
