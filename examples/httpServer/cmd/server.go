package cmd

import (
	"github.com/prismedic/scalpel/config"
	"github.com/spf13/cobra"

	"examples/httpServer/internal/app"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run the API server",
	Long:  `Run the API server.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig(cfgFile)
		app.New().Run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
