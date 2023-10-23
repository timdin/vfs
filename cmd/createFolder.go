package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initCreateFolder(rootCmd *cobra.Command, storage storage.Storage) {
	createFolder := &cobra.Command{
		Use:   constants.CreateFolder,
		Short: "Create folder",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			userName, folderName := args[0], args[1]
			description := ""
			if len(args) == 3 {
				description = args[2]
			}
			err := storage.CreateFolder(userName, folderName, description)
			if err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(createFolder)
}
