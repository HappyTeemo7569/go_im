package client

import (
	"fmt"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"go_im/define"
	"golang.org/x/net/websocket"
	"time"
)

func StartClient() {

	origin := fmt.Sprintf("http://%s:%d", define.Host, define.Port)
	println(origin)
	url := fmt.Sprintf("ws://%s:%d/%s", define.Host, define.Port, define.WebSocketName)
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		tlog.Error(err.Error())
	}

	go func() {
		for {
			msg := []byte("b&h")
			_, _ = ws.Write(msg)
			time.Sleep(10 * time.Second)
		}

	}()

	go func() {
		for {
			var msg = make([]byte, 10240)
			m, err := ws.Read(msg)
			if err != nil {
				tlog.Error(err)
				return
			}
			tlog.Debug("m=%d,Receive: %s", m, msg[:m])
		}
	}()

}

//发送消息
func SendMessage() bool {

	return true
}

//发送心跳
func HeartBeat() {

}
