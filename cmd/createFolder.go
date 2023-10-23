package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

var createFolder = &cobra.Command{
	Use:   constants.CreateFolder,
	Short: "Create folder",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		userName, folderName := args[0], args[1]
		description := ""
		if len(args) == 3 {
			description = args[2]
		}
		err := storage.StorageInstance.CreateFolder(userName, folderName, description)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createFolder)
}
