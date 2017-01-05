package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"gobot.io/x/gobot/platforms/raspi"
)

type Server struct {
	conn     *Connection
	upgrader websocket.Upgrader
	commands map[string]func(*Connection)
}

type Command struct {
	Message string
	Callback func(*Connection)
}

func NewServer() *Server {
	s := &Server{}
	s.upgrader = websocket.Upgrader{}
	s.commands = make(map[string]func(*Connection))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	// Creates websocket connection
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	// Creates device connection
	raspi := raspi.NewAdaptor()
	err = raspi.Connect()
	if err != nil {
		log.Fatal(err)
	}

	s.conn = NewConnection(conn, raspi)

	for {
		_, message, err := s.conn.Read()
		if err != nil {
			log.Println(err)
			break
		}

		if callback, ok := s.commands[message]; ok {
			go callback(s.conn)
		} else {
			log.Println("unregistered command:", message)
		}
	}
}

func (s *Server) RegisterCommand(c *Command) {
	s.commands[c.Message] = c.Callback
}