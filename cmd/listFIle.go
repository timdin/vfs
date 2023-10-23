package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initListFile(rootCmd *cobra.Command, storage storage.Storage) {
	listFile := &cobra.Command{
		Use:   constants.ListFile,
		Short: "List files",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	rootCmd.AddCommand(listFile)
}
