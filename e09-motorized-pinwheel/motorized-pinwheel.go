package main

import "machine"

var (
	switchPin = machine.D2
	motorPin  = machine.D9
)

func main() {
	motorPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	switchPin.Configure(machine.PinConfig{Mode: machine.PinInput})
	for {
		switchState := switchPin.Get()
		if switchState {
			motorPin.High()
		} else {
			motorPin.Low()
		}
	}
}
