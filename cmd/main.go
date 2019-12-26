package main

import piblink "github.com/levsthings/pi-blink"

func main() {
	for i := 0; i < 10; i++ {
		piblink.Blink(26)
	}
}
