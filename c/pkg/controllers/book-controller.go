package controllers

import (
	"c/pkg/models"
	"c/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	NewBooks := models.GetBooks()
	res, _ := json.Marshal(NewBooks)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBook(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	book_id := vars["ID"]
	ID, err := strconv.ParseInt(book_id, 0, 0)
	if err != nil {
		fmt.Println("ID parse error")
	}
	book, _ := models.GetBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["ID"], 0, 0)
	if err != nil {
		fmt.Println("book not exist")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.ParseInt(vars["ID"], 0, 0)
	if err != nil {
		fmt.Println("book not exist")
	}
	var data = &models.Book{}
	utils.ParseBody(r, data)
	b, db := models.GetBook(ID)
	if data.Name != "" {
		b.Name = data.Name
	}
	if data.Author != "" {
		b.Author = data.Author
	}
	if data.Publication != "" {
		b.Publication = data.Publication
	}
	db.Save(&b)
	res, _ := json.Marshal(b)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
