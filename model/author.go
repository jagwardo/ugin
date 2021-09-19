package model

import (
	"github.com/jinzhu/gorm"
)

type AuthorList struct {
	TotalData    int64
	FilteredData int64
	Data         []Author
}

type Author struct {
	gorm.Model
	FirstName string `json:"FirstName" gorm:"type:varchar(255)"`
	LastName  string `json:"LastName" gorm:"type:varchar(255)"`
	Image     string `json:"Image" gorm:"type:varchar(255)"`
}
