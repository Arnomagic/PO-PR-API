package eservice

import (
	"hexapi/databases"
)

type Credential struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type CredentialResponse struct {
	Id    int      `json:"members_id"`
	Rigth []string `json:"rigth"`
}
type AuthEservice interface {
	GetUserByCredentail(Credential) (*CredentialResponse, error)
	AddRigthToId(CredentialResponse) ([]string, error)
	GetRigthByid(int) ([]string, error)
	RemoveRigthByIdFromIndex(int, int) error
}
type authEservice struct {
	db databases.AuthDb
}
