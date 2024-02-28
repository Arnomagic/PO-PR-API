package databases

import "github.com/jmoiron/sqlx"

type Book struct {
	BookID           int    `db:"books_id"`
	Title            string `db:"title"`
	Author           string `db:"author"`
	Publication_year string `db:"publication_year"`
	Genre            string `db:"genre"`
	MemberID         int    `db:"members_id"`
}
type BookDb interface {
	InsertBook(Book) (*Book, error)
	SelectAllBook() ([]Book, error)
	SelectByIdBook(int) (*Book, error)
	UpdateByIdBook(Book) (*Book, error)
}
type bookDb struct {
	db *sqlx.DB
}
