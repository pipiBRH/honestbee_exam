package main

import (
	"github.com/golang/glog"
	_ "honestbee_exam/src/cmd"
	"honestbee_exam/src/socket_server"
)

func main() {
	defer glog.Flush()
	ws.Run()
}
