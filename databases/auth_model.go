package databases

import "github.com/jmoiron/sqlx"

type Credential struct {
	Username string `db:"email"`
}

type CredentialResponse struct {
	Id       int     `db:"members_id"`
	Password string  `db:"password"`
	Rigth    *string `db:"rigth"`
}
type Rigth struct {
	Id    int      `db:"members_id"`
	Rigth []string `db:"rigth"`
}

type AuthDb interface {
	SelectUserByCredentail(Credential) (*CredentialResponse, error)
	InserRigthToId(Rigth) ([]string, error)
	SelectRigthById(int) ([]string, error)
	DeleteRigthByIdFromIndex(int, int) error
}
type authdb struct {
	db *sqlx.DB
}
