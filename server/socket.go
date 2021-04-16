package server

import (
	"encoding/json"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"go_im/define"
	"golang.org/x/net/websocket"
	"time"
)

func WS_Thread(ws *websocket.Conn) {
	Normalchan <- WsMsgItem{CONN_TYPE_CONNECT, ws, nil}

	var err error
	for {
		var msg string

		if err = websocket.Message.Receive(ws, &msg); err != nil {
			tlog.Error("接收消息错误：", err)
			Normalchan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
			return
		}
		//vlog.Debug("接收一般消息", msg)
		skt := GetSocketItem(ws)
		if skt == nil {
			WS_SendError(define.ERROR_TYPE_AUTH, "未保存的连接", ws)
			tlog.Error("未保存的连接")
			Normalchan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
			return
		}

		var WSdata define.ClientMessage
		if err = json.Unmarshal([]byte(msg), &WSdata); err != nil {
			str, _ := json.Marshal(WSdata)
			tlog.Error("解析消息错误：", err, string(str))
			Normalchan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
			return
		}

		skt.UserData.UserId = WSdata.UserId

		if WSdata.MessageType == define.HEARTBEAT_MSG_ID { //心跳包
			skt.MsgTime = int(time.Now().Unix()) //更新时间戳
			sendMsg := &define.TagServer_HeartBeat{
				define.HEARTBEAT_MSG_ID,
			}
			WS_SendMessage(sendMsg, ws)
		} else {
			str, _ := json.Marshal(WSdata)
			tlog.Info("业务逻辑消息：", string(str))
			Normalchan <- WsMsgItem{CONN_TYPE_READ, ws, &WSdata}
		}

	}
}

func WS_OnConnect(ws *websocket.Conn) {
	skt := GetSocketItem(ws)
	if skt != nil {
		tlog.Warning("重复连接...", ws)
	} else {
		AddSocketItem(ws)
	}
}

func WS_Close(ws *websocket.Conn) {

	tlog.Warning("WS_Close开始:", ws)

	ws.Close()
	skt := GetSocketItem(ws)
	if skt == nil {
		tlog.Warning("断开连接，但是没找到连接...", ws)
		return
	}

	DelSocketItem(ws)

	tlog.Warning("WS_Close结束:", ws)
}
