package piblink

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/warthog618/gpio"
)

func Blink(pn int) {
	if err := gpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer gpio.Close()

	pin := gpio.NewPin(pn)
	pin.Output()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go cleanup(c)

	for {
		pin.Toggle()
		time.Sleep(time.Second)
	}
}

func cleanup(c chan os.Signal) {
	<-c

	halt()
}

func halt() {
	gpio.Close()
	os.Exit(1)
}
