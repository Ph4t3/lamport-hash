package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset user password",
	Long:  `Use this command to reset your user password.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reset called")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
