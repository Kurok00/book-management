package main

import (
	"book_management/database"
	"book_management/models"
	"book_management/routes"
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
	database.DB.AutoMigrate(&models.Book{})

	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	book := models.Book{Title: "Test Book", Author: "Test Author", Year: 2023}
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
    // Kết nối database test
    database.Connect()
    database.DB.AutoMigrate(&models.Book{})

    // Tạo một sách mẫu trước
    book := models.Book{Title: "Test Book", Author: "Test Author", Year: 2023}
    database.DB.Create(&book)

    // Cấu hình router
    gin.SetMode(gin.TestMode)
    router := routes.SetupRouter()

    updatedBook := models.Book{Title: "Updated Book", Author: "Updated Author", Year: 2024}
    body, _ := json.Marshal(updatedBook)

    // Ép kiểu book.ID từ uint sang int
    req, _ := http.NewRequest("PUT", "/books/"+strconv.Itoa(int(book.ID)), bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusOK, rr.Code)

    var resultBook models.Book
    err := json.Unmarshal(rr.Body.Bytes(), &resultBook)
    assert.NoError(t, err)
    assert.Equal(t, updatedBook.Title, resultBook.Title)
}