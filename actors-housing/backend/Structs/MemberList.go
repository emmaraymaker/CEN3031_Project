package backen

import (
	"errors"
	"fmt"
)

type Member struct {
	// Union ID
	UID string
	// Name - First, Last
	Name string
	// Phone - ###-###-####
	phone string
	// address@site
	email string
	// currently no minimum requirements
	password string

	isLister bool
	isTenant bool
}

type MemberList struct {
	Ml []Member
}

func (re *MemberList) AddMember(u string, n string, p string, e string, ps string, l bool, t bool) {
	nm := Member{
		UID:      u,
		Name:     n,
		phone:    p,
		email:    e,
		password: ps,
		isLister: l,
		isTenant: t,
	}

	// Add duplicate check

	re.Ml = append(re.Ml, nm)
	fmt.Println(len(re.Ml))
}

func (re MemberList) SearchMember(uid string) (Member, error) {
	for i := 0; i < len(re.Ml); i++ {
		if re.Ml[i].UID == uid {
			return re.Ml[i], nil
		}
	}
	return Member{}, errors.New("member not found")
}

func (re *MemberList) RemoveMember(uid string) {
	for i := 0; i <= len(re.Ml); i++ {
		if re.Ml[i].UID == uid {
			re.Ml = append(re.Ml[:i], re.Ml[i+1:]...)
			return
		}
	}
}
