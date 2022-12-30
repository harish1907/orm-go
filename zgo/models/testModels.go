package models

import (
	"gorm.io/gorm"
)

type State struct {
	gorm.Model
	Name    string    `gorm:"unique" json:"state_name"`
	MyUser  []MyUser  `gorm:"foreignkey:StateID"`
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
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Email     string      `gorm:"unique" json:"email"`
	Password  string      
	StateID   uint        `gorm:"required"`
	VillageID uint        `gorm:"required"`
	Language  []*Language `gorm:"many2many:user_languages;" json:"-"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*MyUser `gorm:"many2many:user_languages;" json:"-"`
}

type Author struct {
	gorm.Model
	Name string `gorm:"unique;not null;"`
	Price uint `gorm:"default:0"`
}

type Award struct {
	gorm.Model
	Name string `gorm:"unique"`
	AuthorID uint 
	Author Author
}