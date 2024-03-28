package internal

import (
	"human_epoch/msg"
	"human_epoch/user"
)

type Room struct {
	Id     int32
	Player [MaxPlayerNumber]*PlayerInfo
}

// 给所有人发消息
func (r *Room) AddRoom(u *user.User) {
	chair := r.MakeSeat()
	r.Player[chair] = &PlayerInfo{
		Anent:      nil,
		Name:       u.Name,
		HeadUrl:    u.HeadUrl,
		State:      1,
		Id:         u.Id,
		EnemyCount: EnemyCount{},
	}
	r.SendAllMsg(&msg.AddRoomRes{
		RoomId: r.Id,
		User:   u,
	})
}

// 分配座位
func (r *Room) MakeSeat() int32 {

	//按顺序坐
	for i := int32(0); i < MaxPlayerNumber; i++ {
		if r.Player[i] == nil {
			return i
		}
	}

	return INVALID_CHAIR
}

// 给所有人发消息
func (r *Room) SendAllMsg(msg interface{}) {
	for _, info := range r.Player {
		if info == nil {
			continue
		}
		info.Anent.WriteMsg(msg)
	}
}

// 给除了某个玩家其他的玩家发消息
func (r *Room) SendOtherOnePlayerMsg(msg interface{}, id int32) {
	for _, info := range r.Player {
		if info == nil || id == info.Id {
			continue
		}
		info.Anent.WriteMsg(msg)
	}
}

// 给单个玩家其他的玩家发消息
func (r *Room) SendOnePlayerMsg(msg interface{}, id int32) {
	if r.Player[id] == nil {
		return
	}
	r.Player[id].Anent.WriteMsg(msg)
}
