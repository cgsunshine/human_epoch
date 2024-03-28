package internal

import "github.com/name5566/leaf/gate"

// PlayerInfo 玩家信息
type PlayerInfo struct {
	Anent   gate.Agent
	Name    string
	HeadUrl string
	State   int32 //玩家状态 0 离线 1 在线 2 游戏中
	Id      int32
	EnemyCount
}

// EnemyCount 灭敌数量
type EnemyCount struct {
	SmallAircraft int
	Boss          int
}
