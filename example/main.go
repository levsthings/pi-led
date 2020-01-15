package main

import (
	piled "github.com/levsthings/pi-led"
	"log"
)

func main() {
	var (
		pin      = 26
		duration = 20 // Seconds
	)
	err := piled.Blink(pin, duration)
	if err != nil {
		log.Fatal("error communicating with the device")
	}
}
