package msg

import (
	"github.com/name5566/leaf/network/json"
	"human_epoch/user"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&SyncPosition{})
	Processor.Register(&Login{})
	Processor.Register(&CreateRoom{})
}

// 登录
type Login struct {
	Id int32
}

// 登录返回
type LoginRes struct {
	Id      int32
	Name    string
	HeadUrl string
}

// 同步位置
type SyncPosition struct {
	X int32
	Y int32
}

// 创建房间
type CreateRoom struct {
	Id int32
}

type CreateRoomRes struct {
	RoomId int32
}

// 加入房间
type AddRoom struct {
	RoomId int32
}

type AddRoomRes struct {
	RoomId int32
	*user.User
}
