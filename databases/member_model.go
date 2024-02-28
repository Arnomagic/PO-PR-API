package databases

import "github.com/jmoiron/sqlx"

type Member struct {
	MemberID int    `db:"members_id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	JoinDate string `db:"join_date"`
	Password string `db:"password"`
}
type MemberDb interface {
	InsertMember(Member) (*Member, error)
	SelectAllMember() ([]Member, error)
	SelectByIdMember(int) (*Member, error)
	UpdateByIdMember(Member) (*Member, error)
}
type memberDb struct {
	db *sqlx.DB
}
