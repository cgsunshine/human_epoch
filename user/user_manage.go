package user

import (
	"github.com/name5566/leaf/gate"
	"sync"
)

var UserManage = new(Users)

type Users struct {
	User  sync.Map //map[id]User
	Agent sync.Map //map[gate.Agent]user
}

func (u *Users) GetUserId(id int32) *User {
	result, ok := u.User.Load(id)
	if !ok {
		return nil
	}
	return result.(*User)
}

func (u *Users) GetUserAgent(agent gate.Agent) *User {
	result, ok := u.User.Load(agent)
	if !ok {
		return nil
	}
	return result.(*User)
}
