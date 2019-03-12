package ws

import (
	"encoding/json"
	"honestbee_exam/src/search"
	"strings"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Sender: "System", Content: "A new socket has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Sender: "System", Content: "A socket has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

type Client struct {
	id     string
	socket *websocket.Conn
	send   chan []byte
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			glog.Errorf("Client read error => %+v\n", err)
			manager.unregister <- c
			c.socket.Close()
			break
		}

		msg := string(message)
		payload := &Message{Sender: c.id, Content: msg}
		if strings.HasPrefix(msg, "@") {
			queryData, err := search.EsQuery(strings.Replace(msg, "@", "", 1))
			if err != nil {
				glog.Errorf("WS payload error => %+v\n", err)
				payload.Content = "query err"
			} else {
				payload.Goods = queryData
			}
		}

		jsonMessage, _ := json.Marshal(payload)
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

type Message struct {
	Sender  string           `json:"sender,omitempty"`
	Content string           `json:"content,omitempty"`
	Goods   []search.HitList `json:"goods,omitempty"`
}

var manager ClientManager

func NewClientManager() ClientManager {
	return ClientManager{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}
