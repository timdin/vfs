package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myvfs",
	Short: "A sample VFS CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please provide a subcommand.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
