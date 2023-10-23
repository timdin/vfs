package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initListFolder(rootCmd *cobra.Command, storage storage.Storage) {
	listFolder := &cobra.Command{
		Use:   constants.ListFolder,
		Short: "List folders",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	rootCmd.AddCommand(listFolder)
}
