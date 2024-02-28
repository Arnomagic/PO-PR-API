package eservice

import "hexapi/databases"

type Loan struct {
	LoansID    int    `db:"loans_id"`
	BooksID    int    `db:"books_id"`
	MembersID  int    `db:"member_id"`
	LoansDate  string `db:"loans_date"`
	ReturnDate string `db:"return_date"`
}
type LoanInsert struct {
	BooksID    int    `db:"books_id"`
	MembersID  int    `db:"member_id"`
	LoansDate  string `db:"loans_date"`
	ReturnDate string `db:"return_date"`
}

type LoanEservice interface {
	AddLoan(LoanInsert) (*Loan, error)
	GetAllLoan() ([]Loan, error)
	GetByIdLoan(int) (*Loan, error)
	EditLoanById(Loan) (*Loan, error)
}
type loanEservice struct {
	db databases.LoanDb
}
