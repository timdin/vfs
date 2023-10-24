package storage

import (
	"github.com/timdin/vfs/configs"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/model"
)

//go:generate mockgen -destination=../mock/storage_mock.go -package=mock github.com/timdin/vfs/storage Storage
type Storage interface {
	Register(nameName string) error
	CreateFolder(userName, folderName, description string) error
	CreateFile(userName, folderName, fileName, description string) error
	DeleteFolder(userName, folderName string) error
	DeleteFile(userName, folderName, fileName string) error
	ListFolder(userName string, sortBy constants.SortByField, order constants.Order) ([]*model.Folder, error)
	ListFile(userName, folderName string, sortBy constants.SortByField, order constants.Order) ([]*model.File, error)
	// RenameFolder(user, folderName, newName string) error
}

func InitStorage() Storage {
	var StorageInstance Storage

	switch configs.AppConfig.Mode {
	case configs.Prod:
		var err error
		StorageInstance, err = initDBProd(configs.AppConfig.DBconfig.Conn)
		if err != nil {
			panic(err)
		}
	case configs.Dev:
		var err error
		StorageInstance, err = initDBDev(configs.AppConfig.DBconfig.Conn)
		if err != nil {
			panic(err)
		}
	}
	return StorageInstance
}
