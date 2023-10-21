package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
)

var deleteFolder = &cobra.Command{
	Use:   constants.DeleteFolder,
	Short: "Delete folder",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(deleteFolder)
}
