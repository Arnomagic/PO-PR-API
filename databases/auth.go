package databases

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hexapi/logs"

	"github.com/jmoiron/sqlx"
)

func NewAuthDb(db *sqlx.DB) AuthDb {
	return authdb{db: db}
}

func (d authdb) SelectUserByCredentail(cred Credential) (*CredentialResponse, error) {
	query := `SELECT m.member_id,
	md.info ->>'password' AS password,
	md.info ->>'rigth' AS rigth 
	FROM library_system.members m 
	JOIN library_system.members_data md ON md.member_id = m.member_id 
	WHERE m.email = $1`
	user := CredentialResponse{}
	err := d.db.Get(&user, query, cred.Username)
	if err != nil {
		logs.Log.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, ErrNoRows
		}
		return nil, ErrDB
	}
	return &user, nil
}

type isRigvalue struct {
	RightVal *string `db:"rigth"`
}

func (d authdb) InserRigthToId(rigth Rigth) ([]string, error) {
	query := "SELECT info->'rigth' AS rigth FROM libraly_system.members_data WHERE member_id = $1 ;"
	isRigval := isRigvalue{}
	err := d.db.Get(&isRigval, query, rigth.Id)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	rigthStr, err := json.Marshal(rigth.Rigth)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	if isRigval.RightVal == nil {
		query = fmt.Sprintf(`UPDATE libraly_system.members_data SET
		info = (info::jsonb || '{"rigth": %v}')
		WHERE member_id = $1`, string(rigthStr))

	} else {
		query = fmt.Sprintf(`UPDATE libraly_system.members_data SET
		info = JSONB_SET(info,'{rigth}',info ->'rigth' || '%v')
		WHERE member_id  = $1`, string(rigthStr))
	}
	res, err := d.db.Exec(query, rigth.Id)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	if i, _ := res.RowsAffected(); i == 0 {
		logs.Log.Error(ErrNoRows.Error())
		return nil, ErrDB
	}
	return rigth.Rigth, nil
}
func (d authdb) SelectRigthById(id int) ([]string, error) {
	query := "SELECT info->'rigth' AS rigth FROM libraly_system.members_data WHERE member_id = $1"
	rigth := isRigvalue{}
	err := d.db.Get(&rigth, query, id)
	if err != nil {
		logs.Log.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, ErrNoRows
		}
		return nil, ErrDB
	}
	if rigth.RightVal == nil {
		return nil, ErrNoRows
	}
	rigthResponse := []string{}
	err = json.Unmarshal([]byte(*rigth.RightVal), &rigthResponse)
	if err != nil {
		return nil, ErrDB
	}
	return rigthResponse, nil
}
func (d authdb) DeleteRigthByIdFromIndex(id int, index int) error {
	query := fmt.Sprintf(`UPDATE libraly_system.members_data SET
	info = JSONB_SET(info,'{rigth}',info->'rigth' #- '{%v}') WHERE member_id = $1
	`, index)

	res, err := d.db.Exec(query, id)
	if err != nil {
		logs.Log.Error(err.Error())
		return ErrDB
	}

	if i, _ := res.RowsAffected(); i == 0 {
		return ErrDB
	}

	return nil
}
