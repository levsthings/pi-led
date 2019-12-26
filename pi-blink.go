package piblink

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/warthog618/gpio"
)

// Blink expects an integer value and you should provide the GPIO pin number
// that you want to control the LED with.
func Blink(p int) {
	if err := gpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer gpio.Close()

	pin := gpio.NewPin(p)
	pin.Output()
	pin.Mode()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go cleanup(c, pin)

	for {
		pin.Toggle()
		time.Sleep(time.Second)
	}

}

func cleanup(c chan os.Signal, p *gpio.Pin) {
	<-c

	p.Low()
	gpio.Close()
	os.Exit(0)
}
