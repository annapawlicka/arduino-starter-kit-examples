package main

import (
	"machine"
	"tinygo.org/x/drivers/tone"
)

var (
	// the numbers below correspond to the frequencies of middle C, D, E, and F
	notes      = [4]uint64{262, 294, 330, 349}
	pwm        = machine.Timer1
	speakerPin = machine.D9
)

func frequencyToPeriod(freq uint64) uint64 {
	return uint64(1e9) / freq
}

func main() {
	// pin that is getting readings from the button
	keyboardPin := machine.ADC{machine.ADC0}
	machine.InitADC()

	// piezo
	piezo, err := tone.New(pwm, speakerPin)
	if err != nil {
		println("could not configure the piezo")
	}

	for {
		keyboardValue := keyboardPin.Get()
		if keyboardValue > 65000 {
			period := frequencyToPeriod(notes[0])
			piezo.SetPeriod(period)
		} else if keyboardValue > 33000 && keyboardValue < 65000 {
			period := frequencyToPeriod(notes[1])
			piezo.SetPeriod(period)
		} else if keyboardValue > 600 && keyboardValue < 33000 {
			period := frequencyToPeriod(notes[2])
			piezo.SetPeriod(period)
		} else if keyboardValue > 5 && keyboardValue < 600 {
			period := frequencyToPeriod(notes[3])
			piezo.SetPeriod(period)
		} else {
			piezo.SetPeriod(0)
		}
	}
}
