package cmd

import (
	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initRegister(rootCmd *cobra.Command, storage storage.Storage) {
	registerCmd := &cobra.Command{
		Use:   constants.Register,
		Short: "Register user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			err := storage.Register(name)
			if err != nil {
				return err
			}
			return nil
		},
	}
	rootCmd.AddCommand(registerCmd)
}
