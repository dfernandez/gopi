package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func service(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)
		if (string(message) == "button_timer_on") {
			go func() {
				time.Sleep(5 * time.Second)
				conn.WriteMessage(websocket.TextMessage, []byte("button_timer_off"))
				log.Print("sent: button_timer_off")
			}()
		}
	}
}

func main() {
	home := http.FileServer(http.Dir("./public"))

	http.Handle("/", home)
	http.HandleFunc("/ws", service)
	http.ListenAndServe(":8080", nil)
}
