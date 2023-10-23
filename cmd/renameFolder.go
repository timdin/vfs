package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initRenameFolder(rootCmd *cobra.Command, storage storage.Storage) {
	renameFolder := &cobra.Command{
		Use:   constants.RenameFolder,
		Short: "Rename folder",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	rootCmd.AddCommand(renameFolder)
}
