package main

import (
	"machine"
	"time"
)

var (
	switchPin       = machine.D9
	switchState     bool
	prevSwitchState bool
	ledIdx                = 0
	previousTime    int64 = 0
	interval        int64 = 600000 // 10 minutes in ms
	leds                  = [6]machine.Pin{machine.D2, machine.D3, machine.D4,
		machine.D5, machine.D6, machine.D7}
)

func main() {
	// set the LED pins as outputs
	for _, led := range leds {
		led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
	// set the tilt switch pin as input
	switchPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	switchState = switchPin.Get()
	prevSwitchState = switchState

	for {
		currentTime := time.Now().UnixMilli()
		if (currentTime-previousTime > interval) && (ledIdx < 6) {
			previousTime = currentTime
			// turn the LED on
			leds[ledIdx].High()
			// increment the led variable
			// in 10 minutes the next LED will light up
			ledIdx++
		}
		// read the switch state
		switchState = switchPin.Get()
		// if the switch has changed
		if switchState != prevSwitchState {
			for _, led := range leds {
				// turn all the LEDs low
				led.Low()
			}
			// reset the led index and the timer
			ledIdx = 0
			previousTime = currentTime
			// set the previous switch state to the current state
			prevSwitchState = switchState
		}
	}
}
