package databases

import (
	"database/sql"
	"fmt"
	"hexapi/logs"

	"github.com/jmoiron/sqlx"
)

func NewMemberDatabases(db *sqlx.DB) MemberDb {
	return memberDb{db: db}
}

func (d memberDb) InsertMember(m Member) (*Member, error) {
	query := "INSERT INTO libraly_system.members (name,email,join_date) VALUES ($1,$2,$3) RETURNING * ;"
	member := Member{}

	trx, err := d.db.Beginx()
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	err = trx.Get(&member, query, m.Name, m.Email, m.JoinDate)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	query = "INSERT INTO libraly_system.members_data (members_id, info) VALUES ($1, $2)"
	res, err := trx.Exec(query, member.MemberID, fmt.Sprintf(`'{"password":"%v"}'`, m.Password))
	if err != nil {
		trx.Rollback()
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	if i, _ := res.RowsAffected(); i == 0 {
		trx.Rollback()
		return nil, ErrDB
	}
	trx.Commit()
	return &member, nil
}
func (d memberDb) SelectAllMember() ([]Member, error) {
	query := "SELECT * FROM libraly_system.members ;"
	members := []Member{}
	err := d.db.Select(&members, query)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	if len(members) == 0 {
		return nil, ErrNoRows
	}
	return members, nil
}
func (d memberDb) SelectByIdMember(id int) (*Member, error) {
	query := "SELECT * FROM libraly_system.members WHERE members_id = $1;"
	members := Member{}
	err := d.db.Get(&members, query, id)
	if err != nil && err != sql.ErrNoRows {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	if err == sql.ErrNoRows {
		return nil, ErrNoRows
	}
	return &members, nil
}
func (d memberDb) UpdateByIdMember(m Member) (*Member, error) {
	fields_values := []struct {
		f string
		v any
	}{}
	if m.Name != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"name", m.Name})
	}
	if m.Email != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"email", m.Email})
	}
	if m.JoinDate != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"join_date", m.JoinDate})
	}
	query := "UPDATE libraly_system.members SET "
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
	query += field + " WHERE members_id = $" + fmt.Sprint(len(argValue)+1) + " RETURNING * ;"
	argValue = append(argValue, m.MemberID)
	member := Member{}
	err := d.db.Get(&member, query, argValue...)
	if err != nil {
		logs.Log.Error(err.Error())
		return nil, ErrDB
	}
	return &member, nil
}
