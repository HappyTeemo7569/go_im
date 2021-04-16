package server

import (
	"encoding/json"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"go_im/define"
	"golang.org/x/net/websocket"
)

//单发消息
func WS_SendMessage(sendMsg interface{}, ws *websocket.Conn) bool {
	msg, err := json.Marshal(sendMsg)
	if err != nil {
		tlog.Error("格式化消息错误：", err)
		WS_SendError(-1, "异常，请重新登录", ws)
		return false
	}

	if err = websocket.Message.Send(ws, string(msg)); err != nil {
		tlog.Error("发送消息错误：", err)
		return false
	}
	tlog.Info("回复消息:", string(msg))
	return true
}

//错误消息
func WS_SendError(code int, msg string, ws *websocket.Conn) bool {
	WSdata := &define.ResultError{
		define.SERVER_MSG_ID_ERROR,
		code,
		msg,
	}
	return WS_SendMessage(WSdata, ws)
}
