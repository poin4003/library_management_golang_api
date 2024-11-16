package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"library_management/global"
	"library_management/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	bookController := NewBookController(global.PDB)
	bookRoutes := r.Group("/books")
	{
		bookRoutes.POST("/", bookController.CreateBook)
		bookRoutes.GET("/", bookController.GetBooks)
	}
	return r
}

func TestCreateBook(t *testing.T) {
	r := setupRouter()

	bookInput := models.BookInput{
		Content:    "Test Book Content",
		Author:     "Test Author",
		PageNumber: 100,
	}

	bookInputJSON, err := json.Marshal(bookInput)
	if err != nil {
		t.Fatalf("Failed to marshal book input: %v", err)
	}

	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(bookInputJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Contains(t, response, "book")
}

func TestGetBooks(t *testing.T) {
	r := setupRouter()

	bookInput := models.BookInput{
		Content:    "Test Book for GET",
		Author:     "Test Author",
		PageNumber: 200,
	}

	bookInputJson, err := json.Marshal(bookInput)
	if err != nil {
		t.Fatalf("could not marshal bookInput: %v", err)
	}

	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(bookInputJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req, _ = http.NewRequest("GET", "/books/", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	books, ok := response["books"].([]interface{})
	assert.True(t, ok)
	assert.Greater(t, len(books), 0, "Books array should not be empty")
}
