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
			userName, folderName := args[0], args[1]
			if err := storage.DeleteFolder(userName, folderName); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(deleteFolder)
}
