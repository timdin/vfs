package storage

import (
	"errors"
	"fmt"

	"github.com/timdin/vfs/constants"
	"github.com/timdin/vfs/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBImpl struct {
	db *gorm.DB
}

func initRemoteDB(conn string) gorm.Dialector {
	return mysql.Open(conn)
}

func initLocalDB(path string) gorm.Dialector {
	return sqlite.Open(path)
}

func initDBProdConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
}

func initDBDevConfig() *gorm.Config {
	return &gorm.Config{}
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
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return err
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
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return err
	}
	existingFolder := &model.Folder{}
	if err := db.lookUpFolder(existingUser, folderName, existingFolder); err != nil {
		return err
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
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return err
	}
	existingFolder := &model.Folder{}
	if err := db.lookUpFolder(existingUser, folderName, existingFolder); err != nil {
		return err
	}
	data := &model.Folder{
		Model: gorm.Model{ID: existingFolder.ID},
	}
	tx := db.db.Begin()
	if err := tx.Delete(data).Error; err != nil {
		// Handle the error, which can occur if the name conflicts
		tx.Rollback()
		return errors.New("Failed to delete folder: " + err.Error())
	}
	if err := tx.Delete(&model.File{}, "folder_id =?", existingFolder.ID).Error; err != nil {
		tx.Rollback()
		return errors.New("Failed to delete files under deleting folder: " + err.Error())
	}
	tx.Commit()
	return nil
}

func (db *DBImpl) DeleteFile(userName, folderName, fileName string) error {
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return err
	}
	existingFolder := &model.Folder{}
	if err := db.lookUpFolder(existingUser, folderName, existingFolder); err != nil {
		return err
	}
	existingFile := &model.File{}
	if err := db.lookUpFile(existingUser, existingFolder, fileName, existingFile); err != nil {
		return err
	}
	data := &model.File{
		Model: gorm.Model{ID: existingFile.ID},
	}

	if err := db.db.Delete(data).Error; err != nil {
		// Handle the error, which can occur if the name conflicts
		return errors.New("Failed to delete file: " + err.Error())
	}
	return nil
}

func (db *DBImpl) ListFolder(userName string, sortBy constants.SortByField, order constants.Order) ([]*model.Folder, error) {
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return nil, err
	}
	var folders []*model.Folder
	if err := db.db.Order(fmt.Sprintf("%s %s", sortBy, order)).Where("user_id =?", existingUser.ID).Find(&folders).Error; err != nil {
		return nil, errors.New("Failed to list folders: " + err.Error())
	}
	return folders, nil
}

func (db *DBImpl) ListFile(userName, folderName string, sortBy constants.SortByField, order constants.Order) ([]*model.File, error) {
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return nil, err
	}
	existingFolder := &model.Folder{}
	if err := db.lookUpFolder(existingUser, folderName, existingFolder); err != nil {
		return nil, err
	}
	var files []*model.File
	if err := db.db.Order(fmt.Sprintf("%s %s", sortBy, order)).Where("user_id =? and folder_id =?", existingUser.ID, existingFolder.ID).Find(&files).Error; err != nil {
		return nil, errors.New("Failed to list files: " + err.Error())
	}
	return files, nil
}

func (db *DBImpl) RenameFolder(userName, folderName, newName string) error {
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return err
	}
	existingFolder := &model.Folder{}
	if err := db.lookUpFolder(existingUser, folderName, existingFolder); err != nil {
		return err
	}
	existingFolder.Name = newName
	if err := db.db.Save(existingFolder).Error; err != nil {
		return err
	}
	return nil
}

func (db *DBImpl) RenameFile(userName, folderName, fileName, newName string) error {
	existingUser := &model.User{}
	if err := db.lookUpUser(userName, existingUser); err != nil {
		return err
	}
	existingFolder := &model.Folder{}
	if err := db.lookUpFolder(existingUser, folderName, existingFolder); err != nil {
		return err
	}
	existingFile := &model.File{}
	if err := db.lookUpFile(existingUser, existingFolder, fileName, existingFile); err != nil {
		return err
	}
	existingFile.Name = newName
	if err := db.db.Save(existingFile).Error; err != nil {
		return err
	}
	return nil
}

func (db *DBImpl) lookUpFile(existingUser *model.User, existingFolder *model.Folder, fileName string, existingFile *model.File) error {
	if err := db.db.Where("user_id =? and folder_id=? and name =?", existingUser.ID, existingFolder.ID, fileName).First(&existingFile).Error; err != nil {
		return fmt.Errorf("File [%s] not found", fileName)
	}
	return nil
}

func (db *DBImpl) lookUpFolder(existingUser *model.User, folderName string, existingFolder *model.Folder) error {
	if err := db.db.Where("user_id =? and name =?", existingUser.ID, folderName).First(&existingFolder).Error; err != nil {
		return fmt.Errorf("Folder [%s] not found", folderName)
	}
	return nil
}

func (db *DBImpl) lookUpUser(userName string, existingUser *model.User) error {
	if err := db.db.Where("name =?", userName).First(&existingUser).Error; err != nil {
		return fmt.Errorf("User [%s] not found", userName)
	}
	return nil
}
