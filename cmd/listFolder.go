package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
)

var listFolder = &cobra.Command{
	Use:   constants.ListFolder,
	Short: "List folders",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(listFolder)
}
