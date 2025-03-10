package cmd

import (
	"devops-terminal/internal/monitor"

	"github.com/spf13/cobra"
)

// monitorCmd represents the monitoring command
var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor system resources (CPU, RAM, disk usage)",
	Run: func(cmd *cobra.Command, args []string) {
		monitor.GetSystemStats() // Call function from internal package
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
}
