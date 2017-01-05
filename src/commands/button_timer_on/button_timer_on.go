package button_timer_on

import (
	"log"
	"time"

	"github.com/dfernandez/gopi/src/server"
)

func NewButtonTimerOn() *server.Command {
	return &server.Command{
		Message: "button_timer_on",
		Callback: func(conn *server.Connection){
			log.Println("button_timer_on")

			time.Sleep(time.Second * 5)

			response := map[string]interface{}{
				"cmd": "button_timer_off",
			}

			conn.WriteJson(response)
		},
	}
}
