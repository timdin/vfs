package model

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	UserID      uint
	FolderID    uint
	Name        string
	Description string
}
