package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

var registerCmd = &cobra.Command{
	Use:   constants.Register,
	Short: "Register user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		err := storage.StorageInstance.Register(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
