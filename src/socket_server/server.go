package ws

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"honestbee_exam/src/schema"
)

func Run() {
	glog.Info("Server GO!")
	manager = NewClientManager()
	go manager.start()
	http.HandleFunc("/ws", wsService)
	http.HandleFunc("/", hcService)
	http.ListenAndServe(schema.Config.Server.GetBind(), nil)
}

func wsService(res http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		glog.Errorf("WS Upgrader error => %+v\n", err)
		http.NotFound(res, req)
		return
	}
	client := &Client{id: uuid.NewV4().String(), socket: conn, send: make(chan []byte)}

	manager.register <- client

	go client.read()
	go client.write()
}

func hcService(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("ok"))
}
