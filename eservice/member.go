package eservice

import (
	"hexapi/configuration"
	"hexapi/databases"
	"hexapi/logs"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func NewMemberEservice(d databases.MemberDb) MemberEservice {
	return memberEservice{db: d}
}
func (e memberEservice) AddMember(m MemberInsert) (*Member, error) {
	if m.Name == "" || m.Email == "" {
		logs.Log.Error(ErrNoDATAINPUT.Error())
		return nil, ErrNoDATAINPUT
	}
	password := "A123456a"
	if m.Password != "" {
		password = m.Password
	}
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	match := hasLowerCase && hasUpperCase && hasNumber && len(password) >= 8
	if !match {
		return nil, ErrBadPassword
	}
	passHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, ErrProcessInterrup
	}
	member := databases.Member{
		Name:     m.Name,
		Email:    m.Email,
		JoinDate: time.Now().Format(time.DateTime),
		Password: string(passHash),
	}
	response, err := e.db.InsertMember(member)
	if err != nil {
		return nil, err
	}
	emember := Member{
		MemberID: response.MemberID,
		Name:     response.Name,
		Email:    response.Email,
		JoinDate: configuration.DateTimeString(response.JoinDate),
	}
	return &emember, nil
}
func (e memberEservice) GetAllMember() ([]Member, error) {
	res, err := e.db.SelectAllMember()
	if err != nil {
		return nil, err
	}
	members := []Member{}
	for _, row := range res {
		members = append(members, Member{
			MemberID: row.MemberID,
			Name:     row.Name,
			Email:    row.Email,
			JoinDate: configuration.DateTimeString(row.JoinDate),
		})
	}
	return members, nil
}
func (e memberEservice) GetByIdMember(id int) (*Member, error) {
	if id == 0 {
		return nil, ErrNoDATAINPUT
	}
	res, err := e.db.SelectByIdMember(id)
	if err != nil {
		return nil, err
	}
	member := Member{
		MemberID: res.MemberID,
		Name:     res.Name,
		Email:    res.Email,
		JoinDate: configuration.DateTimeString(res.JoinDate),
	}
	return &member, nil
}
func (e memberEservice) EditMemberById(m Member) (*Member, error) {
	if m.MemberID == 0 {
		return nil, ErrNoDATAINPUT
	}
	amountDATA := 0
	member := databases.Member{}
	member.MemberID = m.MemberID
	if m.Name != "" {
		member.Name = m.Name
		amountDATA += 1
	}
	if m.Email != "" {
		member.Email = m.Email
		amountDATA += 1
	}
	if m.JoinDate != "" {
		member.JoinDate = m.JoinDate
		amountDATA += 1
	}
	if amountDATA == 0 {
		return nil, ErrNoDataForUpdate
	}
	res, err := e.db.UpdateByIdMember(member)
	if err != nil {
		return nil, err
	}
	response := Member{
		MemberID: res.MemberID,
		Name:     res.Name,
		Email:    res.Email,
		JoinDate: configuration.DateTimeString(res.JoinDate),
	}
	return &response, nil
}
