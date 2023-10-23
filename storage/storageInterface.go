package storage

import (
	"github.com/timdin/vfs/configs"
)

type Storage interface {
	Register(name string) error
	CreateFolder(userName, folderName, description string) error
	// CreateFile(user, folderName, fileName, description string) error
	// DeleteFolder(user, folderName string) error
	// DeleteFile(user, folderName, fileName string) error
	// ListFolder(user, sortBy, order string) error
	// ListFile(user, folderName, sortBy, order string) error
}

var StorageInstance Storage

func InitStorage() {

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

}
