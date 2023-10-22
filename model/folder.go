package model

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	UserID      uint
	Name        string
	Description string
}
