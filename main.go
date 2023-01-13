package main

import (
	"github.com/nlm/briq-cli/commands"
	"github.com/nlm/briq-cli/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.SetEnvPrefix("briq")    // prefix for environment variables
	viper.AutomaticEnv()
	cobra.CheckErr(viper.ReadInConfig())
	cobra.CheckErr(viper.Unmarshal(&config.ConfigFile))
	commands.RootCmd.Execute()
}
