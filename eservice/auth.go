package eservice

import (
	"encoding/json"
	"fmt"
	"hexapi/databases"
	"hexapi/logs"

	"golang.org/x/crypto/bcrypt"
)

func NewAuthEservice(db databases.AuthDb) AuthEservice {
	return authEservice{db: db}
}

func (e authEservice) GetUserByCredentail(cred Credential) (*CredentialResponse, error) {
	if cred.Username == "" || cred.Password == "" {
		return nil, ErrNoDATAINPUT
	}
	credDb := databases.Credential{
		Username: cred.Username,
	}
	res, err := e.db.SelectUserByCredentail(credDb)
	if err != nil {
		return nil, err
	}
	rigth := []string{}
	fmt.Println(res.Rigth)
	if res.Rigth != nil {
		err = json.Unmarshal([]byte(*res.Rigth), &rigth)
		if err != nil {
			logs.Log.Error(err.Error())
			return nil, ErrProcessInterrup
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(cred.Password))
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, databases.ErrNoRows
	}
	credential := CredentialResponse{
		Id:    res.Id,
		Rigth: rigth,
	}
	return &credential, nil
}
func (e authEservice) AddRigthToId(rigth CredentialResponse) ([]string, error) {
	if rigth.Id == 0 || len(rigth.Rigth) == 0 {
		return nil, ErrNoDATAINPUT
	}
	rigthDb := databases.Rigth{
		Id:    rigth.Id,
		Rigth: rigth.Rigth,
	}
	res, err := e.db.InserRigthToId(rigthDb)
	if err != nil {
		return nil, err
	}
	return res, err
}
func (e authEservice) GetRigthByid(id int) ([]string, error) {
	if id == 0 {
		return nil, ErrNoDATAINPUT
	}
	res, err := e.db.SelectRigthById(id)
	if err != nil {
		return nil, ErrNoDATAINPUT
	}
	return res, nil
}
func (e authEservice) RemoveRigthByIdFromIndex(id int, index int) error {
	if id == 0 {
		return ErrNoDATAINPUT
	}
	err := e.db.DeleteRigthByIdFromIndex(id, index)
	if err != nil {
		return err
	}
	return nil
}
