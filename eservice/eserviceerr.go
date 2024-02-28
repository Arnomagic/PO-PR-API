package eservice

import "errors"

var (
	ErrNoDATAINPUT     = errors.New("no input data please again.")
	ErrNoDataForUpdate = errors.New("no data for update please again.")
	ErrProcessInterrup = errors.New("process interrup.")
	ErrBadPassword     = errors.New("password not secure.")
)
