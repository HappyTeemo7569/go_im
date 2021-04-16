package main

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"go_im/client"
	"go_im/server"
)

func main() {
	tlog.Info("start")
	go server.StartServer()
	go client.StartClient()

	select {}
}
