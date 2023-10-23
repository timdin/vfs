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
			return nil
		},
	}
	rootCmd.AddCommand(createFile)
}
