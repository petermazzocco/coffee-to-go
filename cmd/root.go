/*
Copyright © 2025 Peter Mazzocco <petermazzocco@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

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
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
