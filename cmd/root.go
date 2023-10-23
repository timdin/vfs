package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/storage"
)

func InitCmd(storage storage.Storage) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "myvfs",
		Short: "A sample VFS CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please provide a subcommand.")
		},
	}
	initRegister(rootCmd, storage)
	initCreateFile(rootCmd, storage)
	initCreateFolder(rootCmd, storage)

	return rootCmd
}
