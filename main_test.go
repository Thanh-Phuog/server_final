package main

import (
	"book_mana/database"
	"book_mana/models"
	"book_mana/routes"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	// Kết nối database test
	database.Connect()
	//database.DB.AutoMigrate(&models.Book{})

	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	book := models.Book{Title: "Test Book", Author: "Test Author", Year: "2023"}
	body, _ := json.Marshal(book)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var createdBook models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &createdBook)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, createdBook.Title)
}

func TestGetBooks(t *testing.T) {
	// Kết nối database test
	database.Connect()

	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdateBook(t *testing.T) {
	database.Connect()
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	book := models.Book{Title: "Book", Author: "Test Author", Year: "2024"}
	database.DB.Create(&book)

	book.Title = "Updated Book"
	body, _ := json.Marshal(book)
	// Sử dụng ID của sách thay vì tiêu đề để cập nhật
	req, _ := http.NewRequest("PUT", "/books/"+strconv.Itoa(int(book.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var updatedBook models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &updatedBook)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, updatedBook.Title)
}
func TestDeleteBook(t *testing.T) {
	database.Connect()
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()
	// Tạo một cuốn sách để xóa
	book := models.Book{Title: "Book to Delete", Author: "Test Author", Year: "2024"}
	database.DB.Create(&book)

	req, _ := http.NewRequest("DELETE", "/books/"+strconv.Itoa(int(book.ID)), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestSearchBook(t *testing.T) {
	database.Connect()
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	database.DB.Create(&models.Book{Title: "Book 1", Author: "Test Author", Year: "2024"})
	database.DB.Create(&models.Book{Title: "Book 2", Author: "Test Author", Year: "2025"})
	database.DB.Create(&models.Book{Title: "Another Book 2", Author: "Test Author", Year: "2026"})
	database.DB.Create(&models.Book{Title: "Book 3", Author: "Test Author", Year: "2026"})

	req, _ := http.NewRequest("GET", "/books/search?title=2", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var books []models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &books)
	assert.NoError(t, err)

	for _, book := range books {
		assert.Contains(t, book.Title, "Book 2")
	}
}
