package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookCreation(t *testing.T) {
	book := Book{
		Content:    "Test Book",
		Author:     "Test Author",
		PageNumber: 100,
	}

	assert.NotNil(t, book.ID)
	assert.NotNil(t, book.CreatedAt)
	assert.NotNil(t, book.UpdatedAt)
}

func TestBookUpdate(t *testing.T) {
	book := Book{
		Content:    "Test Book",
		Author:     "Test Author",
		PageNumber: 100,
	}

	book.Content = "Updated Title"
	assert.NotNil(t, book.ID)
	assert.NotNil(t, book.CreatedAt)
	assert.NotNil(t, book.UpdatedAt)
}
