# pi-led


Utility for blinking LEDs via a Raspberry Pi. 


### Details


This project uses a GPIO pin for controlling an LED.


### Usage


You can import `pi-led` as a library, and blink an LED for a given duration.

```go
package main

import (
	piled "github.com/levsthings/pi-led"
)

func main() {
	var (
		pin      = 26
		duration = 20 // Seconds
	)
	piled.Blink(pin, duration)
}
```

The above example also lives in the `example` folder which can be built and tested.
