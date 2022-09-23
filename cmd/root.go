package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var configFile string

// init will be called whenever a package initialize in golang
func init() {
	cobra.OnInitialize(initConfig)
	// Developers are going to define flags and configuration settings here
	RootCmd.PersistentFlags().StringVar(
		&configFile,
		"config",
		"",
		"Absolute path to configuration file (default at $HOME directory)",
	)
}

// RootCmd represent for the base command if it called without any subcommand
var RootCmd = &cobra.Command{
	Use:   "go-service",
	Short: "Boilerplate project of backend service",
	Long:  "Boilerplate project of backend service adheres to System Design and Architecture Guidelines and DDD Design",
	Run: func(cmd *cobra.Command, args []string) {
		// Associate codes
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}

// initConfig Setup default values and load configuration file
func initConfig() {
	// Init configuration values
}
