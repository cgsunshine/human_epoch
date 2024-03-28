package internal

import (
	"human_epoch/msg"
	"math/rand"
)

var RoomManger map[int32]*Room

func init() {
	RoomManger = make(map[int32]*Room, MapInitCap)
}

// 给所有人发消息
func CreateRoom(c *msg.CreateRoom) int32 {
	room := new(Room)
	id := int32(rand.Intn(9999-1000+1) + 1000)
	room.Id = id
	RoomManger[id] = room
	return id
}
