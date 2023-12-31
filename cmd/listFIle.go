package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initListFile(rootCmd *cobra.Command, storage storage.Storage) {
	sortByNameArg, sortByCreatedArg := false, false

	listFile := &cobra.Command{
		Use:   constants.ListFile,
		Short: "List files",
		Long: `Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|
		desc]`,
		Example: `list-files [username] [foldername] [--sort-name|--sort-created] [asc|
		desc]`,
		Args: cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, folder := args[0], args[1]
			// set default sort by field to name
			applySortField := constants.SortByName
			if sortByCreatedArg {
				applySortField = constants.SortByCreated
			}
			// set default order to asc
			applyOrder := constants.OrderAsc
			if len(args) > 2 {
				orderArg := args[2]
				if order, err := constants.ConstructOrder(orderArg); err != nil {
					return err
				} else {
					applyOrder = order
				}
			}
			res, err := storage.ListFile(user, folder, applySortField, applyOrder)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				fmt.Printf("Warning: The folder %s/%s is empty\n", user, folder)
			} else {
				for _, file := range res {
					fmt.Printf("%s\t%s\t%s\n", file, folder, user)
				}
			}
			return nil
		},
	}
	listFile.Flags().BoolVarP(&sortByCreatedArg, constants.SortByCreatedFlag, "C", false, "Sort by created")
	listFile.Flags().BoolVarP(&sortByNameArg, constants.SortByNameFlag, "N", false, "Sort by name")
	rootCmd.AddCommand(listFile)
}
