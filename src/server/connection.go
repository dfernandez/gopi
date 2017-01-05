package server

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"gobot.io/x/gobot/platforms/raspi"
)

type Connection struct {
	Raspi *raspi.Adaptor
	conn *websocket.Conn
	read sync.Mutex
	write sync.Mutex
}

func NewConnection(conn *websocket.Conn, raspi *raspi.Adaptor) *Connection {
	return &Connection{
		Raspi: raspi,
		conn: conn,
		read: sync.Mutex{},
		write: sync.Mutex{},
	}
}

func (c *Connection) Read() (int, string, error) {
	c.read.Lock()
	tp, message, err := c.conn.ReadMessage()
	c.read.Unlock()

	return tp, string(message), err
}

func (c *Connection) WriteJson(message map[string]interface{}) {
	jsonResponse, _ := json.Marshal(message)

	c.write.Lock()
	c.conn.WriteMessage(websocket.TextMessage, jsonResponse)
	c.write.Unlock()
}
