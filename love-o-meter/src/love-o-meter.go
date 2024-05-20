package main

import (
	"machine"
	"time"
)

// room temperature in Celsius
var baselineTemp = 20.0

func main() {
	// temp sensor (TMP36, analog sensor)
	tempSensor := machine.ADC{machine.ADC0}
	machine.InitADC()

	// LEDs
	led1 := machine.D2
	led2 := machine.D3
	led3 := machine.D4
	led1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led1.Low()
	led2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led2.Low()
	led3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led3.Low()

	var maxValue = 65535.0

	for {
		// Get() returns a value that is always scaled in the range [0..65535],
		// regardless of the ADC’s configured bit size (resolution)
		sensorValue := tempSensor.Get()
		// convert the ADC reading to voltage
		sensorValueAsFloat := float64(sensorValue)
		voltage := (sensorValueAsFloat / maxValue) * 5.0
		// convert to Celsius
		temperature := (voltage - 0.5) * 100.0

		if temperature < baselineTemp {
			led1.Low()
			led2.Low()
			led3.Low()
		} else if (temperature >= baselineTemp) && (temperature < baselineTemp+3) {
			led1.High()
			led2.Low()
			led3.Low()
		} else if temperature >= (baselineTemp + 5) {
			led1.High()
			led2.High()
			led3.High()
		}
		time.Sleep(1 * time.Second)
	}
}
