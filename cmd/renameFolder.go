package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
)

var renameFolder = &cobra.Command{
	Use:   constants.RenameFolder,
	Short: "Rename folder",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(renameFolder)
}
