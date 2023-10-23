package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initCreateFile(rootCmd *cobra.Command, storage storage.Storage) {
	createFile := &cobra.Command{
		Use:   constants.CreateFile,
		Short: "Create file",
		Args:  cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			userName, folderName, fileName := args[0], args[1], args[2]
			description := ""
			if len(args) == 4 {
				description = args[3]
			}
			if err := storage.CreateFile(userName, folderName, fileName, description); err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(createFile)
}
