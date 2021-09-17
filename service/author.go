package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yakuter/ugin/model"
)

func GetAuthor(db *gorm.DB, id string) (*model.Author, error) {
	var err error
	author := new(model.Author)

	if err := db.Where("id = ? ", id).First(&author).Error; err != nil {
		log.Println(err)

		return nil, err
	}

	return author, err
}

func GetAuthors(c *gin.Context, db *gorm.DB, args model.Args) ([]model.Author, int64, int64, error) {
	authors := []model.Author{}
	var filteredData, totalData int64

	table := "authors"
	query := db.Select(table + ".*")
	query = query.Offset(Offset(args.Offset))
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))

	if err := query.Find(&authors).Error; err != nil {
		log.Println(err)
		return authors, filteredData, totalData, err
	}

	// // Count filtered table
	// // We are resetting offset to 0 to return total number.
	// // This is a fix for Gorm offset issue
	query = query.Offset(0)
	query.Table(table).Count(&filteredData)

	// // Count total table
	db.Table(table).Count(&totalData)

	return authors, filteredData, totalData, nil
}

// SavePost both cretes and updates post according to if ID field is empty or not
func SaveAuthor(db *gorm.DB, author *model.Author) (*model.Author, error) {
	if err := db.Save(&author).Error; err != nil {
		return author, err
	}

	return author, nil
}

// DeletePost soft deletes all records.
func DeleteAuthor(db *gorm.DB, id string) error {
	author := new(model.Author)
	if err := db.Where("id = ? ", id).Delete(&author).Error; err != nil {
		log.Println(err)
		return err
	}

	tag := new(model.Tag)
	if err := db.Where("author_id = ? ", id).Delete(&tag).Error; err != nil {
		log.Println(err)
	}

	return nil
}
