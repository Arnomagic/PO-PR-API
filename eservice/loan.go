package eservice

import (
	"fmt"
	"hexapi/databases"
)

func NewLoanEservice(d databases.LoanDb) LoanEservice {
	return loanEservice{db: d}
}

func (d loanEservice) AddLoan(l LoanInsert) (*Loan, error) {
	if l.LoansDate == "" || l.ReturnDate == "" {
		fmt.Println("please input data !")
		return nil, ErrNoDATAINPUT
	}
	loan := databases.Loan{
		LoansDate:  l.LoansDate,
		ReturnDate: l.ReturnDate,
	}
	response, err := d.db.InsertLoan(loan)
	if err != nil {
		return nil, err
	}
	eloan := Loan{
		LoansID:    response.LoansID,
		BooksID:    response.BooksID,
		MembersID:  response.MembersID,
		LoansDate:  response.LoansDate,
		ReturnDate: response.ReturnDate,
	}
	return &eloan, nil
}
func (d loanEservice) GetAllLoan() ([]Loan, error) {
	res, err := d.db.SelectAllLoan()
	if err != nil {
		return nil, err
	}
	loan := []Loan{}
	for _, row := range res {
		loan = append(loan, Loan{
			LoansID:    row.LoansID,
			BooksID:    row.BooksID,
			MembersID:  row.MembersID,
			LoansDate:  row.LoansDate,
			ReturnDate: row.ReturnDate,
		})
	}
	return loan, nil
}
func (d loanEservice) GetByIdLoan(id int) (*Loan, error) {
	if id == 0 {
		return nil, ErrNoDATAINPUT
	}
	res, err := d.db.SelectByIdLoan(id)
	if err != nil {
		return nil, err
	}
	loan := Loan{
		LoansID:    res.LoansID,
		BooksID:    res.BooksID,
		MembersID:  res.MembersID,
		LoansDate:  res.LoansDate,
		ReturnDate: res.ReturnDate,
	}
	return &loan, nil
}
func (d loanEservice) EditLoanById(l Loan) (*Loan, error) {
	if l.LoansID == 0 {
		return nil, ErrNoDATAINPUT
	}
	amountDATA := 0
	loan := databases.Loan{}
	loan.LoansID = l.LoansID
	if l.LoansDate != "" {
		loan.LoansDate = l.LoansDate
		amountDATA += 1
	}
	if l.ReturnDate != "" {
		loan.ReturnDate = l.ReturnDate
		amountDATA += 1
	}
	res, err := d.db.UpdateByIdLoan(loan)
	if err != nil {
		return nil, err
	}
	response := Loan{
		LoansID:    res.LoansID,
		BooksID:    res.BooksID,
		MembersID:  res.MembersID,
		LoansDate:  res.LoansDate,
		ReturnDate: res.ReturnDate,
	}
	return &response, nil
}
