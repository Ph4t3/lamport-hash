package cmd

import (
	"l-hash-frontend/handlers"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your account",
	Long:  `Use this command to Login to your user account using Lamport Hash`,
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
