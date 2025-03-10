package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clipboardCmd represents the multiple clipboard storage feature
var clipboardCmd = &cobra.Command{
	Use:   "clipboard",
	Short: "Store and manage multiple clipboard entries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Clipboard management tool... (Feature under development)")
	},
}

func init() {
	rootCmd.AddCommand(clipboardCmd)
}
