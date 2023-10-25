package model

import (
	"fmt"

	"github.com/timdin/vfs/constants"
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	UserID      uint
	Name        string
	Description string
}

func (f *Folder) BeforeSave(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(f).Where("name = ? and user_id = ?", f.Name, f.UserID).Count(&count)
	if count > 0 {
		return fmt.Errorf("Folder with the name [%s] already exists", f.Name)
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
	return fmt.Sprintf("%v\t%v\t%v", f.Name, d, f.Model.CreatedAt.Format(constants.TimeFormat))
}
