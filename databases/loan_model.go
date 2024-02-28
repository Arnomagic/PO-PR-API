package databases

import "github.com/jmoiron/sqlx"

type Loan struct {
	LoansID    int    `db:"loans_id"`
	BooksID    int    `db:"books_id"`
	MembersID  int    `db:"member_id"`
	LoansDate  string `db:"loans_date"`
	ReturnDate string `db:"return_date"`
}
type LoanDb interface {
	InsertLoan(Loan) (*Loan, error)
	SelectAllLoan() ([]Loan, error)
	SelectByIdLoan(int) (*Loan, error)
	UpdateByIdLoan(Loan) (*Loan, error)
}
type loanDb struct {
	db *sqlx.DB
}
