package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initDeleteFolder(rootCmd *cobra.Command, storage storage.Storage) {
	deleteFolder := &cobra.Command{
		Use:   constants.DeleteFolder,
		Short: "Delete folder",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	rootCmd.AddCommand(deleteFolder)
}
