package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"library_management/models"
	"net/http"
)

type cBook struct {
	db *gorm.DB
}

func NewBookController(db *gorm.DB) *cBook {
	return &cBook{db: db}
}

func (c *cBook) CreateBook(ctx *gin.Context) {
	var bookInput models.BookInput

	if err := ctx.ShouldBindJSON(&bookInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Content:    bookInput.Content,
		Author:     bookInput.Author,
		PageNumber: bookInput.PageNumber,
	}

	if err := c.db.WithContext(ctx).
		Create(&book).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

func (c *cBook) GetBooks(ctx *gin.Context) {
	var books []models.Book
	c.db.WithContext(ctx).
		Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"books": books})
}

func (c *cBook) GetBook(ctx *gin.Context) {
	var book models.Book

	id := ctx.Param("id")

	if err := c.db.WithContext(ctx).
		First(&book, id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"post": nil})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"post": nil})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

func (c *cBook) UpdateBook(ctx *gin.Context) {
	var bookInput models.BookInput
	var bookFound models.Book

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&bookInput); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.db.WithContext(ctx).
		First(&bookFound, id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"post": nil})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"post": nil})
		return
	}

	bookFound.Content = bookInput.Content
	bookFound.Author = bookInput.Author
	bookFound.PageNumber = bookInput.PageNumber

	if err := c.db.WithContext(ctx).
		Updates(&bookFound).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"post": nil})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "update success"})
}

func (c *cBook) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var post models.Book

	if err := c.db.WithContext(ctx).
		First(&post, id).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"post": nil})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"post": nil})
		return
	}

	if err := c.db.WithContext(ctx).
		Delete(&post, id).
		Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"post": nil})
	}

	ctx.Status(http.StatusNoContent)
}
