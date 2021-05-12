package cmd

import (
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "l-hash-frontend",
	Short: "A CLI application for Lamport Hash",
	Long:  `A CLI application build using GO for the demonstration of Lamport Hash.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
