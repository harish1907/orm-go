package models

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type State struct {
	gorm.Model
	Name    string    `gorm:"unique" json:"state_name"`
	MyUser  []MyUser  `gorm:"foreignkey:StateID" json:"-"`
	Village []Village `gorm:"foreignkey:StateID"`
}

type Village struct {
	gorm.Model
	Name    string   `gorm:"unique"  json:"village_name"`
	MyUser  []MyUser `gorm:"foreignkey:VillageID" json:"-"`
	StateID uint     `json:"-"`
}

type MyUser struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"`
	StateID   uint   `gorm:"required"`
	VillageID uint   `gorm:"required"`
	Language  []*Language `gorm:"many2many:user_languages;" json:"-"`
}

func (c *MyUser) PasswordHash() error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(c.Password), 10)
	c.Password = string(hash)
	return nil
}

type Language struct {
	gorm.Model
	Name  string
	Users []*MyUser `gorm:"many2many:user_languages;" json:"-"`
}
