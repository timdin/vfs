package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
)

var registerCmd = &cobra.Command{
	Use:   constants.Register,
	Short: "Register user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg1 := args[0]
		fmt.Printf("Register: %s\n", arg1)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}
