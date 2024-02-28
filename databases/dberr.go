package databases

import "errors"

var (
	ErrDB     = errors.New("database error")
	ErrNoRows = errors.New("no row in db")
)
