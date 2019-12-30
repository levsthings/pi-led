# pi-blink


Utility for blinking LEDs via a Raspberry Pi. 


### Details


This project uses a GPIO for controlling an LED.


### Usage


You can import `pi-blink` as a library, and blink an LED for a given duration.

```go
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
```

The above example also lives in the `example` folder which can be built and tested.
