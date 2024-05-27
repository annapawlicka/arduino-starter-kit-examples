package main

import (
	"annapawlicka.arduino-fun/tools"
	"machine"
	"time"
	"tinygo.org/x/drivers/tone"
)

var (
	ledPin                = machine.LED
	sensorLow      uint16 = 1023
	sensorHigh     uint16 = 0
	pwm                   = machine.Timer1
	speakerPin            = machine.D9 // PB1
	sensorValue    uint16
	frequencyValue uint16
	period         uint64
)

func main() {
	ledPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	lightSensor := machine.ADC{machine.ADC0}
	machine.InitADC()
	// configure piezo (speaker)
	speaker, err := tone.New(pwm, speakerPin)
	if err != nil {
		println("could not configure the speaker")
	}

	// calibrate for the first 5 seconds after program runs
	ledPin.High()
	calibrationTime := time.Now().Add(5 * time.Second)
	for time.Now().Before(calibrationTime) {
		sensorValue = lightSensor.Get()
		if sensorValue > sensorHigh {
			sensorHigh = sensorValue
		}
		if sensorValue < sensorLow {
			sensorLow = sensorValue
		}
	}
	// turn the LED off, signaling the end of the calibration period
	ledPin.Low()

	for {
		sensorValue = lightSensor.Get()
		// map the sensor values to a wide range of pitches
		// Arduino produces frequencies in the range of 50 - 4000
		frequencyValue = tools.ScaleValue(sensorValue, float64(sensorLow), float64(sensorHigh), 50, 4000)
		// calculate period
		if frequencyValue == 0 {
			period = 0 // no sound when the sensor is completely covered
		} else {
			period = uint64(1e9) / uint64(frequencyValue)
		}
		// play the tone for 20 ms on pin 8
		speaker.SetPeriod(period)
		// wait for a moment
		time.Sleep(time.Millisecond * 10)
	}
}
