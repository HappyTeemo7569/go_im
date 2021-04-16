package main

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"go_im/client"
	"go_im/server"
)

func main() {
	tlog.Info("11")
	go server.StartServer()
	go client.StartClient()
}
