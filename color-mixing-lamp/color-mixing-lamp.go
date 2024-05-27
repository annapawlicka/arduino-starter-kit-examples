package main

import (
	"annapawlicka.arduino-fun/tools"
	"machine"
	"time"
)

var scaleMultiplier float64 = 256.0 / 65535.0

// Arduino API's AnalogRead() returns values in (0 - 1024) range.
// In TinyGo API, ADC's Get() method returns a uint16 value in (0-65535) range.
// However, PWM's Set() is 8-bit and so it accepts values in (0 - 256) range.
// Additionally, ADC's Get() returns uint16, but PWM's Set() accepts uint32.
func scaleFromADCToPWM(value uint16) uint32 {
	return uint32(tools.ScaleValue(value, 0, 65535, 0, 256))
}

func main() {
	// GREEN
	// Set up PWM for pin 9. Uses Timer1, channel A.
	pwm1 := machine.Timer1
	greenLEDPin := machine.PB1 // pin 9 on the Uno
	greenLEDchannel, errGreenLEDchannel := pwm1.Channel(greenLEDPin)
	if errGreenLEDchannel != nil {
		println("Error assigning pin to greenLEDchannel:", errGreenLEDchannel.Error())
	}

	// RED
	// Set up PWM for pin 10. Uses the same Timer1, but channel B.
	redLEDPin := machine.PB2 // pin 10 on the Uno
	redLEDchannel, errRedLEDchannel := pwm1.Channel(redLEDPin)
	if errRedLEDchannel != nil {
		println("Error assigning pin to redLEDchannel:", errRedLEDchannel.Error())
	}

	// BLUE
	//Set up PWM for pin 11. Uses Timer2, channel A.
	pmw2 := machine.Timer2
	pmw2.Configure(machine.PWMConfig{})
	blueLEDPin := machine.PB3 // pin 11 on the Uno
	blueLEDchannel, errBlueLEDchannel := pmw2.Channel(blueLEDPin)
	if errBlueLEDchannel != nil {
		println("Error assigning pin to blueLEDchannel:", errBlueLEDchannel.Error())
	}

	// pins with photoresistors
	redSensorPin := machine.ADC{machine.ADC0}
	greenSensorPin := machine.ADC{machine.ADC1}
	blueSensorPin := machine.ADC{machine.ADC2}

	machine.InitADC()

	for {
		// Read the sensors
		redSensorValue := redSensorPin.Get()
		time.Sleep(time.Millisecond * 5)
		greenSensorValue := greenSensorPin.Get()
		time.Sleep(time.Millisecond * 5)
		blueSensorValue := blueSensorPin.Get()

		redValue := scaleFromADCToPWM(redSensorValue)
		greenValue := scaleFromADCToPWM(greenSensorValue)
		blueValue := scaleFromADCToPWM(blueSensorValue)

		pwm1.Set(redLEDchannel, redValue)
		pwm1.Set(greenLEDchannel, greenValue)
		pmw2.Set(blueLEDchannel, blueValue)
	}
}
