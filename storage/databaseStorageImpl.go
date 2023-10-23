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

func (db *DBImpl) CreateFolder(userName, folderName, description string) error {
	// look up the user with the given name, fail if not found
	var existingUser *model.User
	if err := db.db.Where("name =?", userName).First(&existingUser).Error; err != nil {
		return errors.New("Failed to find user: " + err.Error())
	}

	data := &model.Folder{
		Name:        folderName,
		UserID:      existingUser.ID,
		Description: description,
	}
	if err := db.db.Create(data).Error; err != nil {
		// Handle the error, which can occur if the name conflicts
		return errors.New("Failed to create folder: " + err.Error())
	}
	return nil
}

func (db *DBImpl) CreateFile(userName, folderName, fileName, description string) error {
	var existingUser *model.User
	if err := db.db.Where("name =?", userName).First(&existingUser).Error; err != nil {
		return errors.New("Failed to find user: " + err.Error())
	}
	var existingFolder *model.Folder
	if err := db.db.Where("user_id =? and name =?", existingUser.ID, folderName).First(&existingFolder).Error; err != nil {
		return errors.New("Failed to find folder: " + err.Error())
	}
	data := &model.File{
		Name:        fileName,
		UserID:      existingUser.ID,
		FolderID:    existingFolder.ID,
		Description: description,
	}
	if err := db.db.Create(data).Error; err != nil {
		// Handle the error, which can occur if the name conflicts
		return errors.New("Failed to create folder: " + err.Error())
	}
	return nil
}

func (db *DBImpl) DeleteFolder(userName, folderName string) error {
	var existingUser *model.User
	if err := db.db.Where("name =?", userName).First(&existingUser).Error; err != nil {
		return errors.New("Failed to find user: " + err.Error())
	}
	var existingFolder *model.Folder
	if err := db.db.Where("user_id =? and name =?", existingUser.ID, folderName).First(&existingFolder).Error; err != nil {
		return errors.New("Failed to find folder: " + err.Error())
	}
	data := &model.Folder{
		Model: gorm.Model{ID: existingFolder.ID},
	}
	if err := db.db.Delete(data).Error; err != nil {
		// Handle the error, which can occur if the name conflicts
		return errors.New("Failed to delete folder: " + err.Error())
	}
	return nil
}
