package main

import (
	"machine"
	"time"
)

func main() {
	ledGreen := machine.D3
	ledGreen.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledRed1 := machine.D4
	ledRed1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledRed2 := machine.D5
	ledRed2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !button.Get() {
			// button not pressed
			ledGreen.High()
			ledRed1.Low()
			ledRed2.Low()
		} else {
			ledGreen.Low()
			// while button pressed, keep toggling the red LEDs every 250 ms
			ledRed1.Low()
			ledRed2.High()
			time.Sleep(time.Millisecond * 250)
			ledRed1.High()
			ledRed2.Low()
			time.Sleep(time.Millisecond * 250)
		}
	}

}
