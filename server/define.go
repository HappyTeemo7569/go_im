package server

import (
	"go_im/define"
	"golang.org/x/net/websocket"
	"sync"
)

//网络连接
var SocketList sync.Map

//关闭通道
var CloseChan chan int

type SocketItem struct {
	BConnect bool
	UserData *define.UserItem
	MsgTime  int
}

var Normalchan QueChan

type QueChan chan WsMsgItem

type WsMsgItem struct {
	Conntype int
	Ws       *websocket.Conn
	WSdata   *define.ClientMessage
}

//websocket 逻辑
const (
	CONN_TYPE_CONNECT = iota //连接
	CONN_TYPE_READ           //通讯
	CONN_TYPE_CLOSE          //关闭
)
