package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	briqSecretKey = ""
)

var rootCmd = cobra.Command{
	Use: "briq-cli",
}

func Register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func main() {
	briqSecretKey = os.Getenv("BRIQ_SECRET_KEY")
	if briqSecretKey == "" {
		log.Fatal("BRIQ_SECRET_KEY is not defined")
	}
	rootCmd.Execute()
}
