package server

import (
	"fmt"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/astaxie/beego/toolbox"
	"go_im/define"
	"golang.org/x/net/websocket"
	"net/http"
)

func StartServer() {
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				fmt.Printf("%v", e)
			}
		}()
		pattern := fmt.Sprintf("/%s", define.WebSocketName)
		http.Handle(pattern, websocket.Handler(WS_Thread))
		tlog.Info("启动webSocket:", pattern)

		addr := fmt.Sprintf(":%d", define.Port)
		tlog.Info("启动webSocket:监听:", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			tlog.Error("websocket启动失败！", err)
			CloseChan <- 0
		} else {
			tlog.Info("开始监听", addr)
		}
	}()

	//每10秒检测心跳
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				fmt.Printf("%v", e)
			}
		}()
		tk := toolbox.NewTask("CheckSocketList", "0/1 * * * * *", CheckSocketList)
		err := tk.Run()
		if err != nil {
			fmt.Println(err)
		}
		toolbox.AddTask("CheckSocketList", tk)
		toolbox.StartTask()
		tlog.Info("添加定时任务CheckSocketList")
	}()

	<-CloseChan
	tlog.Emergency("收到关闭消息")
}
