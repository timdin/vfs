package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
)

var listFile = &cobra.Command{
	Use:   constants.ListFile,
	Short: "List files",
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(listFile)
}
