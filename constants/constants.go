package constants

import (
	"fmt"
	"strings"
)

const (
	// subcommands
	Register     = "register"
	CreateFolder = "create-folder"
	DeleteFolder = "delete-folder"
	RenameFolder = "rename-folder"
	ListFolder   = "list-folders"
	CreateFile   = "create-file"
	DeleteFile   = "delete-file"
	ListFile     = "list-files"
	// listing flags
	SortByNameFlag    = "sort-name"
	SortByCreatedFlag = "sort-created"
)

const (
	TestDB = "./test_database.db"
)

type SortByField string

const (
	SortByName    SortByField = "name"
	SortByCreated SortByField = "created_at"
)

type Order string

const (
	OrderAsc        Order  = "asc"
	OrderDesc       Order  = "desc"
	OrderAscString  string = "asc"
	OrderDescString string = "desc"
)

func ConstructOrder(order string) (Order, error) {
	switch strings.ToLower(order) {
	case OrderAscString:
		return OrderAsc, nil
	case OrderDescString:
		return OrderDesc, nil
	default:
		return "", fmt.Errorf("invalid order: %s", order)
	}
}
