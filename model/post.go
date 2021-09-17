package model

import (
	"github.com/jinzhu/gorm"
)

// Data is mainle generated for filtering and pagination
type Data struct {
	TotalData    int64
	FilteredData int64
	Data         []Post
}

type Args struct {
	Sort   string
	Order  string
	Offset string
	Limit  string
	Search string
}

type Post struct {
	gorm.Model
	Title       string `json:"Title" gorm:"type:varchar(255)"`
	Description string `json:"Description" gorm:"type:text"`
	Content     string `gorm:"type:longtext" json:"Content"`
	Author      Author `json:"Author" binding:"required" gorm:"foreignkey:AuthorID"`
	AuthorID    uint64
	Tags        []Tag // One-To-Many relationship (has many - use Tag's UserID as foreign key)
}

type Tag struct {
	gorm.Model
	PostID      uint   `gorm:"index"` // Foreign key (belongs to)
	Name        string `json:"Name" gorm:"type:varchar(255)"`
	Description string `json:"Description" gorm:"type:text"`
}

// You can send this data to API /posts/ endpoint with POST method to create dummy content
/*
{
    "Name": "Hello World",
    "Description": "Lorem ipsum  dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
    "Tags": [
        {
            "Name": "lorem",
            "Description": "This is lorem tag"
		},
		{
            "Name": "ipsum",
            "Description": "This is ipsum tag"
        }
    ]
}
*/
