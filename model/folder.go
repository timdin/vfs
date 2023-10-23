package model

import (
	"errors"

	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	UserID      uint
	Name        string
	Description string
}

func (f *Folder) BeforeCreate(tx *gorm.DB) (err error) {
	var existingFolder Folder

	if err := tx.Where("name = ? and user_id = ?", f.Name, f.UserID).First(&existingFolder).Error; err == nil {
		return errors.New("Folder with the same name already exists")
	}
	return nil
}
