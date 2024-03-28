package gate

import (
	"human_epoch/game"
	"human_epoch/login"
	"human_epoch/msg"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&msg.SyncPosition{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.CreateRoom{}, game.ChanRPC)

	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)
}
