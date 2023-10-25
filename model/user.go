package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	var existingUser User
	if err := tx.Where("name = ?", u.Name).First(&existingUser).Error; err == nil {
		return fmt.Errorf("User with the name [%s] already exists", u.Name)
	}
	return nil
}
