package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initRenameFile(rootCmd *cobra.Command, storage storage.Storage) {
	renameFile := &cobra.Command{
		Use:   constants.RenameFile,
		Short: "Rename file",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			userName, folderName, fileName, newFileName := args[0], args[1], args[2], args[3]
			if err := storage.RenameFile(userName, folderName, fileName, newFileName); err != nil {
				return err
			}
			fmt.Printf("File [%s] renamed to [%s] successfully\n", fileName, newFileName)
			return nil
		},
	}
	rootCmd.AddCommand(renameFile)
}
