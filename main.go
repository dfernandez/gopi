package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"encoding/json"
	"runtime"
	"os"
)

var upgrader = websocket.Upgrader{}

var startTime time.Time

func status(conn *websocket.Conn) {
	var m runtime.MemStats
	c := time.Tick(10 * time.Second)
	hostname, _ := os.Hostname()

	for now := range c {
		runtime.ReadMemStats(&m)

		response := map[string]interface{}{
			"cmd": "system_status",
			"now": now,
			"host": hostname,
			"arch": runtime.GOARCH,
			"os": runtime.GOOS,
			"mem": m.TotalAlloc,
			"uptime": time.Since(startTime).String(),
		}
		jsonResponse, _ := json.Marshal(response)
		conn.WriteMessage(websocket.TextMessage, jsonResponse)
	}
}

func service(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	go status(conn)

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
				response := map[string]interface{}{
					"cmd": "button_timer_off",
				}
				jsonResponse, _ := json.Marshal(response)
				conn.WriteMessage(websocket.TextMessage, jsonResponse)
				log.Print("sent: button_timer_off")
			}()
		}
	}
}

func main() {
	startTime = time.Now()
	home := http.FileServer(http.Dir("/home/gopi/public"))

	http.Handle("/", home)
	http.HandleFunc("/ws", service)
	http.ListenAndServe("192.168.1.33:8080", nil)
}
