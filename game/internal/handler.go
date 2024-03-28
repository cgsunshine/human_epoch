package internal

import (
	"github.com/name5566/leaf/gate"
	"human_epoch/msg"
	"human_epoch/user"
	"reflect"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&msg.SyncPosition{}, handleSyncPosition)
	handler(&msg.CreateRoom{}, handleCreateRoom)
	handler(&msg.AddRoom{}, handleAddRoom)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// 同步位置消息
func handleSyncPosition(args []interface{}) {
	m := args[0].(*msg.SyncPosition)
	// 消息的发送者 拿到id
	a := args[1].(gate.Agent)
	u := user.UserManage.GetUserAgent(a)
	RoomManger[u.RoomId].SendOtherOnePlayerMsg(m, u.Id)
}

// 创建房间
func handleCreateRoom(args []interface{}) {
	m := args[0].(*msg.CreateRoom)
	// 消息的发送者 拿到id
	a := args[1].(gate.Agent)
	u := user.UserManage.GetUserAgent(a)
	roomId := CreateRoom(m)
	u.RoomId = roomId
	a.WriteMsg(&msg.CreateRoomRes{RoomId: roomId})
}

// 加入房间
func handleAddRoom(args []interface{}) {
	//m := args[0].(*msg.AddRoom)
	// 消息的发送者 拿到id
	a := args[1].(gate.Agent)
	u := user.UserManage.GetUserAgent(a)
	RoomManger[u.RoomId].AddRoom(u)
}
