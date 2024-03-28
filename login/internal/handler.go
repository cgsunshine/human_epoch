package internal

import (
	"github.com/name5566/leaf/gate"
	"human_epoch/msg"
	"human_epoch/user"
	"reflect"
)

func init() {
	handler(&msg.Login{}, handLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

var userSMap = map[int32]*msg.LoginRes{
	1: {
		Id:      1,
		Name:    "11",
		HeadUrl: "11",
	},
	2: {
		Id:      2,
		Name:    "22",
		HeadUrl: "22",
	},
	3: {
		Id:      3,
		Name:    "33",
		HeadUrl: "33",
	},
	4: {
		Id:      4,
		Name:    "44",
		HeadUrl: "44",
	},
	5: {
		Id:      5,
		Name:    "55",
		HeadUrl: "55",
	},
}

// 登录
func handLogin(args []interface{}) {
	m := args[0].(*msg.Login)
	// 消息的发送者 拿到id
	a := args[1].(gate.Agent)
	result := userSMap[m.Id]
	u := &user.User{
		Id:      result.Id,
		Name:    result.Name,
		HeadUrl: result.HeadUrl,
		RoomId:  0,
	}
	//user add
	user.UserManage.User.Store(m.Id, u)
	user.UserManage.Agent.Store(a, u)
	a.WriteMsg(result)
}
