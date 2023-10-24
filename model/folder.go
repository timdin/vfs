package model

import (
	"errors"
	"fmt"
	"time"

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

// create a stringer for file, but the user information is not included
// these information will be appended in the command handler
func (f *Folder) String() string {
	// set default value for description as N/A
	d := "N/A"
	if f.Description != "" {
		d = f.Description
	}
	return fmt.Sprintf("%v\t%v\t%v", f.Name, d, f.Model.CreatedAt.Format(time.RFC3339))
}
