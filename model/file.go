package model

import (
	"errors"
	"fmt"
	"time"

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

// create a stringer for file, but the user and folder information is not included
// these information will be appended in the command handler
func (f *File) String() string {
	// set default value for description as N/A
	d := "N/A"
	if f.Description != "" {
		d = f.Description
	}
	return fmt.Sprintf("%v\t%v\t%v", f.Name, d, f.Model.CreatedAt.Format(time.RFC3339))
}
