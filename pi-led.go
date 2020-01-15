package piled

import (
	"errors"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/warthog618/gpio"
)

// Blink expects an integer value, you should provide the GPIO pin number
// that you want to control the LED with.
func Blink(p int, s int) error {
	if err := gpio.Open(); err != nil {
		return errors.New("couldn't establish GPIO connection")
	}
	defer gpio.Close()

	pin := gpio.NewPin(p)
	pin.Output()
	pin.Mode()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go cleanup(c, pin)

	for i := 0; i < s; i++ {
		pin.Toggle()
		time.Sleep(time.Second)
	}
	return nil
}

func cleanup(c chan os.Signal, p *gpio.Pin) {
	<-c

	p.Low()
	os.Exit(0)
}
