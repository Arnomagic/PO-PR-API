package eservice

import "hexapi/databases"

type Book struct {
	BookID           int    `db:"books_id"`
	Title            string `db:"title"`
	Author           string `db:"author"`
	Publication_year string `db:"publication_year"`
	Genre            string `db:"genre"`
}
type BookInsert struct {
	Title            string `db:"title"`
	Author           string `db:"author"`
	Publication_year string `db:"publication_year"`
	Genre            string `db:"genre"`
}
type BookEservice interface {
	AddBook(BookInsert) (*Book, error)
	GetAllBook() ([]Book, error)
	GetByIdBook(int) (*Book, error)
	EditBookById(Book) (*Book, error)
}
type bookEservice struct {
	db databases.BookDb
}
