package main

import (
	piblink "github.com/levsthings/pi-blink"
)

func main() {
	var (
		pin      = 26
		duration = 20 // Seconds
	)
	piblink.Blink(pin, duration)
}
