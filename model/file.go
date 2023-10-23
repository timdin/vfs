package model

import (
	"errors"

	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	UserID      uint
	FolderID    uint
	Name        string
	Description string
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	var existingFile File

	if err := tx.Where("name = ? and user_id = ? and folder_id=?", f.Name, f.UserID, f.FolderID).First(&existingFile).Error; err == nil {
		return errors.New("File with the same name under same folder already exists")
	}
	return nil
}
