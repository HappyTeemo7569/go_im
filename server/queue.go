package server

import (
	"encoding/json"
	"github.com/HappyTeemo7569/teemoKit/tlog"
)

//一般队列
func queueNormal() {
	for {
		tlog.Debug("等待一般消息")
		data := <-Normalchan
		msg, _ := json.Marshal(data.WSdata)
		tlog.Debug("收到一般消息", string(msg))

		switch data.Conntype {
		case CONN_TYPE_CONNECT:
			WS_OnConnect(data.Ws)
		case CONN_TYPE_CLOSE:
			WS_Close(data.Ws)
		case CONN_TYPE_READ: //业务逻辑

		}
	}
}
