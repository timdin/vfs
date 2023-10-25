package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/storage"
	"github.com/timdin/vfs/validation"
)

func InitCmd(storage storage.Storage) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "myvfs",
		Short: "A sample VFS CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please provide a subcommand.")
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := validation.InvalidCharacterValidation(args); err != nil {
				return err
			}
			return nil
		},
	}
	initRegister(rootCmd, storage)
	initCreateFile(rootCmd, storage)
	initCreateFolder(rootCmd, storage)
	initDeleteFolder(rootCmd, storage)
	initDeleteFile(rootCmd, storage)
	initListFolder(rootCmd, storage)
	initListFile(rootCmd, storage)
	initRenameFolder(rootCmd, storage)

	return rootCmd
}
