package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initDeleteFile(rootCmd *cobra.Command, storage storage.Storage) {
	deleteFile := &cobra.Command{
		Use:   constants.DeleteFile,
		Short: "Delete file",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			userName, folderName, fileName := args[0], args[1], args[2]
			if err := storage.DeleteFile(userName, folderName, fileName); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(deleteFile)
}
