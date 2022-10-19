package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jplindgren/bookstore/pkg/models"
	"github.com/jplindgren/bookstore/pkg/utils"
)

func GetBooks (w http.ResponseWriter, r *http.Request) {
	allBooks := models.GetAllBooks()
	payload, _ := json.Marshal(allBooks)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func GetBooksById (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error parsing book id: %s", bookId)
	}

	book, _ := models.GetBookById(id)
	payload, _ := json.Marshal(book)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func CreateBook (w http.ResponseWriter, r *http.Request) {
	var bookModel = &models.Book{}
	utils.ParseBody(r, &bookModel)
	newBook := bookModel.CreateBook()
	payload, _ := json.Marshal(newBook)
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func DeleteBook (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error parsing book id: %s", bookId)
	}

	deletedBook := models.DeleteBook(id)

	payload, _ := json.Marshal(deletedBook)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func UpdateBook (w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Printf("Error parsing book id: %s", bookId)
	}

	book, db := models.GetBookById(id)
	utils.ParseBody(r, &updatedBook)
	if (updatedBook.Name != book.Name){
		book.Name = updatedBook.Name
	}
	if (updatedBook.Author != book.Author){
		book.Author = updatedBook.Author
	}
	if (updatedBook.Publication != book.Publication){
		book.Publication = updatedBook.Publication
	}

	db.Save(&book)

	payload, _ := json.Marshal(book)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
