package storage

import (
	"os"

	"github.com/timdin/vfs/configs"
	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/model"
	"gorm.io/gorm"
)

func InitStorage(config *configs.Config) Storage {
	var storageConfig *gorm.Config
	var storageType gorm.Dialector
	var db *gorm.DB
	var err error

	switch config.DBmode {
	case configs.Prod:
		storageConfig = initDBProdConfig()
	case configs.Dev:
		storageConfig = initDBDevConfig()
	}

	switch config.DBtype {
	case configs.RemoteStorage:
		storageType = initRemoteDB(config.RemoteDBconfig.Conn)
	case configs.LocalStorage:
		storageType = initLocalDB(config.LocalDBconfig.Path)
	}

	if db, err = gorm.Open(storageType, storageConfig); err != nil {
		panic(err)
	}
	db = migrateDB(db)

	return &DBImpl{db}
}

func InitTestDB() *DBImpl {
	db, err := gorm.Open(initLocalDB(constants.TestDB), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = migrateDB(db)
	return &DBImpl{
		db: db,
	}
}
func TeardownTestDB() {
	os.Remove(constants.TestDB)
}

func migrateDB(db *gorm.DB) *gorm.DB {
	migrateErr := db.AutoMigrate(model.User{}, model.Folder{}, model.File{})
	if migrateErr != nil {
		panic(migrateErr)
	}
	return db
}
