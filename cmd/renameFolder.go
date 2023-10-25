package cmd

import (
	"fmt"

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
			userName, folderName, newFolderName := args[0], args[1], args[2]
			if err := storage.RenameFolder(userName, folderName, newFolderName); err != nil {
				return err
			}
			fmt.Printf("Folder [%s] renamed to [%s] successfully\n", folderName, newFolderName)
			return nil
		},
	}
	rootCmd.AddCommand(renameFolder)
}
