package button_timer_on

import (
	"log"
	"time"

	"github.com/dfernandez/gopi/src/server"
	"gobot.io/x/gobot/drivers/gpio"
)

func NewButtonTimerOn() *server.Command {
	return &server.Command{
		Message: "button_timer_on",
		Callback: func(conn *server.Connection){
			log.Println("button_timer_on")
			led := gpio.NewLedDriver(conn.Raspi, "11")
			err := led.On()
			if err != nil {
				log.Println(err)
			}

			time.Sleep(time.Second * 5)

			log.Println("button_timer_off")
			err = led.Off()
			if err != nil {
				log.Println(err)
			}

			response := map[string]interface{}{
				"cmd": "button_timer_off",
			}

			conn.WriteJson(response)
		},
	}
}
