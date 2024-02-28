package databases

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func NewLoanDatabases(db *sqlx.DB) LoanDb {
	return loanDb{db: db}
}

func (d loanDb) InsertLoan(l Loan) (*Loan, error) {
	query := "INSERT INTO libraly_system.loans (books_id,member_id,loans_date,return_date) VALUES ($1,$2,$3,$4) RETURNING * ;"
	loan := Loan{}
	err := d.db.Get(&loan, query, l.BooksID, l.MembersID, l.LoansDate, l.ReturnDate)
	if err != nil {
		fmt.Println(err)
		return nil, ErrDB
	}
	return &loan, nil
}
func (d loanDb) SelectAllLoan() ([]Loan, error) {
	query := "SELECT * FROM libraly_system.loans ;"
	loan := []Loan{}
	err := d.db.Select(&loan, query)
	if err != nil {
		fmt.Println(err)
		return nil, ErrDB
	}
	if len(loan) == 0 {
		return nil, ErrNoRows
	}
	return loan, nil
}
func (d loanDb) SelectByIdLoan(id int) (*Loan, error) {
	query := "SELECT * FROM libraly_system.loans WHERE loans_id = $1;"
	loan := Loan{}
	err := d.db.Get(&loan, query, id)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return nil, ErrDB
	}
	if err == sql.ErrNoRows {
		return nil, ErrNoRows
	}
	return &loan, nil
}
func (d loanDb) UpdateByIdLoan(l Loan) (*Loan, error) {
	fields_values := []struct {
		f string
		v any
	}{}
	if l.LoansDate != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"loans_date", l.LoansDate})
	}
	if l.ReturnDate != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"return_date", l.ReturnDate})
	}
	query := "UPDATE libraly_system.loans SET "
	field := ""
	argValue := []any{}
	for i, row := range fields_values {
		if i == 0 {
			field += "" + row.f + " = $" + fmt.Sprint((i + 1))
			argValue = append(argValue, row.v)
		} else {
			field += ", " + row.f + " = $" + fmt.Sprint((i + 1))
			argValue = append(argValue, row.v)
		}
	}
	query += field + " WHERE loans_id = $" + fmt.Sprint(len(argValue)+1) + " RETURNING * ;"
	argValue = append(argValue, l.LoansID)
	fmt.Println(query)
	fmt.Println(argValue...)
	loan := Loan{}
	err := d.db.Get(&loan, query, argValue...)
	if err != nil {
		fmt.Println(err)
		return nil, ErrDB
	}
	return &loan, nil
}
