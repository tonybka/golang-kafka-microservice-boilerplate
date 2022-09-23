package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Command to start running server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.WithField("CLI:command", "serve").Info("Start running server")
		// Start server from here ... (being continued)
	},
}

func init() {
	RootCmd.AddCommand(ServeCmd)
}
