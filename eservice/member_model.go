package eservice

import (
	"hexapi/databases"
)

type Member struct {
	MemberID int    `json:"members_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	JoinDate string `json:"join_date"`
}
type MemberInsert struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MemberEservice interface {
	AddMember(MemberInsert) (*Member, error)
	GetAllMember() ([]Member, error)
	GetByIdMember(int) (*Member, error)
	EditMemberById(Member) (*Member, error)
}
type memberEservice struct {
	db databases.MemberDb
}
