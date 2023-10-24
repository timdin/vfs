package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/storage"
)

func initListFolder(rootCmd *cobra.Command, storage storage.Storage) {
	sortByNameArg, sortByCreatedArg := false, false
	listFolder := &cobra.Command{
		Use:   constants.ListFolder,
		Short: "List folders",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			user := args[0]
			// set default sort by field to name
			applySortField := constants.SortByName
			if sortByCreatedArg {
				applySortField = constants.SortByCreated
			}
			// set default order to asc
			applyOrder := constants.OrderAsc
			if len(args) > 1 {
				orderArg := args[1]
				if order, err := constants.ConstructOrder(orderArg); err != nil {
					return err
				} else {
					applyOrder = order
				}
			}
			res, err := storage.ListFolder(user, applySortField, applyOrder)
			if err != nil {
				return err
			}
			for _, folder := range res {
				fmt.Printf("%s\t%s\n", folder, user)
			}
			return nil
		},
	}
	listFolder.Flags().BoolVarP(&sortByCreatedArg, constants.SortByCreatedFlag, "C", false, "Sort by created")
	listFolder.Flags().BoolVarP(&sortByNameArg, constants.SortByNameFlag, "N", false, "Sort by name")
	rootCmd.AddCommand(listFolder)
}
