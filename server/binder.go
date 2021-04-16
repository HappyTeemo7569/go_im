package server

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"go_im/define"
	"golang.org/x/net/websocket"
	"time"
)

//检测socket连接
func CheckSocketList() error {
	timenow := int(time.Now().Unix())
	//println(timenow)
	SocketList.Range(func(k, v interface{}) bool {
		switch v.(type) {
		case *SocketItem:
			value := v.(*SocketItem)
			if value != nil && (!value.BConnect && timenow-value.MsgTime >= 12) {
				value.BConnect = false
				tlog.Info("超时断开用户，当前时间:%d，用户时间:%d", timenow, value.MsgTime)
				ws := k.(*websocket.Conn)
				Normalchan <- WsMsgItem{CONN_TYPE_CLOSE, ws, nil}
				return true
			}
		}
		return false
	})
	return nil
}

func GetSocketItem(ws *websocket.Conn) *SocketItem {
	var item *SocketItem
	value, _ := SocketList.Load(ws)
	if value != nil {
		item = value.(*SocketItem)
	}
	return item
}

func AddSocketItem(ws *websocket.Conn) {
	SocketList.Store(ws, &SocketItem{
		true,
		&define.UserItem{},
		int(time.Now().Unix()),
	})
}

func DelSocketItem(ws *websocket.Conn) {
	SocketList.Delete(ws)
}
