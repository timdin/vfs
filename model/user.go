package model

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var existingUser User
	if err := tx.Where("name = ?", u.Name).First(&existingUser).Error; err == nil {
		return errors.New("User with the same name already exists")
	}
	return nil
}
