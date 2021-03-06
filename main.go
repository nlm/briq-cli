package main

import (
	"github.com/nlm/briq-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	viperKeyBriqSecretKey = "secret_key"
)

var (
	config  Config
	rootCmd = cobra.Command{
		Use:   "briq-cli",
		Short: "briq command-line utility",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(utils.PrefixError("config error", config.Check()))
		},
	}
)

func Register(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func main() {
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.SetEnvPrefix("briq")    // prefix for environment variables
	viper.AutomaticEnv()
	cobra.CheckErr(viper.ReadInConfig())
	cobra.CheckErr(viper.Unmarshal(&config))
	rootCmd.Execute()
}
