package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dockerCmd represents the docker management command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Manage Docker containers easily",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Docker management tool... (Feature under development)")
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}
