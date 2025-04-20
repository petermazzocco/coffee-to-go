/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "coffee",
	Short: "A simple cli command to learn Golang with Coffee examples",
	Long: `Struggling to learn Golang? Coffee can help!
 			)
         (
      _________
     |         |
     | COFFEE  |
     |  TO GO  |
     |_________|
      \       /
       '-----'

Coffee To Go is a CLI library for Go that explains multiple concepts
and provides practical examples. Simply run the command "coffee" to get started!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
