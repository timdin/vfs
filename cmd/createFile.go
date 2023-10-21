package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
)

var createFile = &cobra.Command{
	Use:   constants.CreateFile,
	Short: "Create file",
	Args:  cobra.RangeArgs(3, 4),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(createFile)
}
