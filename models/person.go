package models

import "github.com/jinzhu/gorm"

// Person struct model
type Person struct {
	gorm.Model
	PrivyID   string `json:"privyid,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Address   string `json:"address,omitempty"`
}
