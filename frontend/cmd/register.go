package cmd

import (
	"l-hash-frontend/handlers"

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register an account",
	Long:  `Use this command to Register a user using Lamport Hash`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Register()
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
