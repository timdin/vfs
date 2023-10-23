package storage

import (
	"github.com/timdin/vfs/configs"
)

//go:generate mockgen -destination=../mock/storage_mock.go -package=mock github.com/timdin/vfs/storage Storage
type Storage interface {
	Register(name string) error
	CreateFolder(userName, folderName, description string) error
	CreateFile(user, folderName, fileName, description string) error
	// DeleteFolder(user, folderName string) error
	// DeleteFile(user, folderName, fileName string) error
	// ListFolder(user, sortBy, order string) error
	// ListFile(user, folderName, sortBy, order string) error
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
