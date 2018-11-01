package models

import "github.com/jinzhu/gorm"

// User struct model
type User struct {
	gorm.Model
	PrivyID   string `json:"privyid,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Address   string `json:"address,omitempty"`
}
