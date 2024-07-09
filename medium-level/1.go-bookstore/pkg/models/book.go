package models

import (
	"my-go-projects/go-bookstore/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Book represents the structure of our resource
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook adds a new book to the database
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById retrieves a book by its ID
func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a book by its ID
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
