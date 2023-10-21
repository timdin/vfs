package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
)

var deleteFile = &cobra.Command{
	Use:   constants.DeleteFile,
	Short: "Delete file",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(deleteFile)
}
