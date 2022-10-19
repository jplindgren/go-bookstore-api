package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/jplindgren/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name 				string `gorm:"json:name"`
	Author 			string `gorm:"json:author"`
	Publication string `gorm:"json:publication"`
}

func init() {
	config.Connect();
	db = config.GetDB()
	test := db.AutoMigrate(Book{});
	if test != nil && test.Error != nil {
		//We have an error
		log.Fatal(fmt.Sprintf("%s with error %s", "Failed", db.Error))
	}
	fmt.Println("Success")
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var allBooks []Book
	db.Find(&allBooks)
	return allBooks;
}

func GetBookById(bookId int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?").Find(&book)
	return &book, db
}

func DeleteBook(bookId int64) Book {
	var book Book
	db.Where("ID=?").Delete(&book)
	return book
}
