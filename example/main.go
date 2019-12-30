package main

import (
	piled "github.com/levsthings/pi-blink"
)

func main() {
	var (
		pin      = 26
		duration = 20 // Seconds
	)
	piled.Blink(pin, duration)
}
