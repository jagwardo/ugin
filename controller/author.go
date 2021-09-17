package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yakuter/ugin/model"
	"github.com/yakuter/ugin/service"
)

// GetAuthor godoc
// @Summary Show an Author
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Author ID"
// @Success 200 "Success"
// @Router /authors/{id} [get]
func (base *Controller) GetAuthor(c *gin.Context) {
	id := c.Params.ByName("id")

	author, err := service.GetAuthor(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
	}

	c.JSON(200, author)
}

// GetAuthors godoc
// @Summary List Authors
// @Accept  json
// @Produce  json
// @Success 200 "Success"
// @Router /authors/ [get]
func (base *Controller) GetAuthors(c *gin.Context) {
	var args model.Args

	// Define and get sorting field
	args.Sort = c.DefaultQuery("Sort", "ID")

	// Define and get sorting order field
	args.Order = c.DefaultQuery("Order", "DESC")

	// Define and get offset for pagination
	args.Offset = c.DefaultQuery("Offset", "0")

	// Define and get limit for pagination
	args.Limit = c.DefaultQuery("Limit", "25")

	// Get search keyword for Search Scope
	args.Search = c.DefaultQuery("Search", "")

	// Fetch results from database
	authors, filteredData, totalData, err := service.GetAuthors(c, base.DB, args)
	if err != nil {
		c.AbortWithStatus(404)
	}

	// Fill return data struct
	data := model.AuthorList{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         authors,
	}

	c.JSON(200, data)
}

// CreatePost godoc
// @Summary Create Post
// @Description Create Post
// @Content Create Post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param Post body object true "Post"
// @Success 200 "Success"
// @Router /posts/ [post]
func (base *Controller) CreateAuthor(c *gin.Context) {
	author := new(model.Author)

	err := c.ShouldBindJSON(&author)

	if err != nil {
		log.Print(err)
		c.AbortWithError(400, err)
		return
	}

	author, err = service.SaveAuthor(base.DB, author)
	if err != nil {
		c.AbortWithError(404, err)
		return
	}

	c.JSON(200, author)
}

// UpdatePost godoc
// @Summary Update Post
// @Description Update Post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Param Post body object true "Post"
// @Success 200 "Success"
// @Router /posts/{id} [put]
func (base *Controller) UpdateAuthor(c *gin.Context) {
	id := c.Params.ByName("id")

	author, err := service.GetAuthor(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	err = c.ShouldBindJSON(&author)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	author, err = service.SaveAuthor(base.DB, author)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, author)
}

// DeletePost godoc
// @Summary Delete Post
// @Description Delete Post
// @Tags posts
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 "Success"
// @Router /posts/{id} [delete]
func (base *Controller) DeleteAuthor(c *gin.Context) {
	id := c.Params.ByName("id")

	err = service.DeleteAuthor(base.DB, id)
	if err != nil {
		c.AbortWithStatus(404)
		return
	}

	c.JSON(200, gin.H{"id#" + id: "deleted"})
}
