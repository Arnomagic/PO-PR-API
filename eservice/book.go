package eservice

import (
	"fmt"
	"hexapi/databases"
)

func NewBookEservice(d databases.BookDb) BookEservice {
	return bookEservice{db: d}
}

func (d bookEservice) AddBook(b BookInsert) (*Book, error) {
	if b.Title == "" || b.Author == "" || b.Publication_year == "" || b.Genre == "" {
		fmt.Println("please input data !")
		return nil, ErrNoDATAINPUT
	}
	book := databases.Book{
		Title:            b.Title,
		Author:           b.Author,
		Publication_year: b.Publication_year,
		Genre:            b.Genre,
	}
	response, err := d.db.InsertBook(book)
	if err != nil {
		return nil, err
	}
	ebook := Book{
		BookID:           response.BookID,
		Title:            response.Title,
		Author:           response.Author,
		Publication_year: response.Publication_year,
		Genre:            response.Genre,
	}
	return &ebook, nil
}
func (d bookEservice) GetAllBook() ([]Book, error) {
	res, err := d.db.SelectAllBook()
	if err != nil {
		return nil, err
	}
	books := []Book{}
	for _, row := range res {
		books = append(books, Book{
			BookID:           row.BookID,
			Title:            row.Title,
			Author:           row.Author,
			Publication_year: row.Publication_year,
			Genre:            row.Genre,
		})
	}
	return books, nil
}
func (d bookEservice) GetByIdBook(id int) (*Book, error) {
	if id == 0 {
		return nil, ErrNoDATAINPUT
	}
	res, err := d.db.SelectByIdBook(id)
	if err != nil {
		return nil, err
	}
	books := Book{
		BookID:           res.BookID,
		Title:            res.Title,
		Author:           res.Author,
		Publication_year: res.Publication_year,
		Genre:            res.Genre,
	}
	return &books, nil
}
func (d bookEservice) EditBookById(b Book) (*Book, error) {
	if b.BookID == 0 {
		return nil, ErrNoDATAINPUT
	}
	amountDATA := 0
	book := databases.Book{}
	book.BookID = b.BookID
	if b.Title != "" {
		book.Title = b.Title
		amountDATA += 1
	}
	if b.Author != "" {
		book.Author = b.Author
		amountDATA += 1
	}
	if b.Publication_year != "" {
		book.Publication_year = b.Publication_year
		amountDATA += 1
	}
	if b.Genre != "" {
		book.Genre = b.Genre
		amountDATA += 1
	}
	if amountDATA == 0 {
		return nil, ErrNoDataForUpdate
	}
	res, err := d.db.UpdateByIdBook(book)
	if err != nil {
		return nil, err
	}
	response := Book{
		BookID:           res.BookID,
		Title:            res.Title,
		Author:           res.Author,
		Publication_year: res.Publication_year,
		Genre:            res.Genre,
	}
	return &response, nil
}
