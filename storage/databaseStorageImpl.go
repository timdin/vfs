package storage

import (
	"errors"

	"github.com/timdin/vfs/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBImpl struct {
	db *gorm.DB
}

func initDBProd(dbConfig string) (*DBImpl, error) {
	return initDB(dbConfig, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}
func initDBDev(dbConfig string) (*DBImpl, error) {
	return initDB(dbConfig, &gorm.Config{})
}

func initDB(dbConfig string, gormConfig *gorm.Config) (*DBImpl, error) {
	db, err := gorm.Open(mysql.Open(dbConfig), gormConfig)
	if err != nil {
		return nil, err
	}
	migrateErr := db.AutoMigrate(model.User{}, model.Folder{}, model.File{})
	if migrateErr != nil {
		return nil, err
	}
	return &DBImpl{
		db: db,
	}, nil
}

func (db *DBImpl) Register(name string) error {
	data := &model.User{
		Name: name,
	}
	if err := db.db.Create(data).Error; err != nil {
		// Handle the error, which can occur if the name conflicts
		return errors.New("Failed to create user: " + err.Error())
	}
	return nil
}
