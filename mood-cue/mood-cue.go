package main

import (
	"annapawlicka.arduino-fun/tools"
	"machine"
	"time"
	"tinygo.org/x/drivers/servo"
)

var (
	pwm = machine.Timer1
	pin = machine.D9
)

func main() {
	s, err := servo.New(pwm, pin)
	if err != nil {
		println("could not configure servo")
	}

	// analog pin used to connect the potentiometer
	potPin := machine.ADC{machine.ADC0}
	machine.InitADC()

	for {
		// read the value of the potentiometer
		potValue := potPin.Get()
		// scale the numbers from the potentiometer
		// ADC's Get() returns values in 0-65535 range, but Servo supports
		// pulse range between 1000-2000µs for 90 degrees rotation
		// 0 degrees is 1000, 45 degrees is 1500 and 90 degrees is 2000.
		angle := tools.ScaleValue(potValue, 0, 65535, 1000, 2000)
		s.SetMicroseconds(int16(angle))
		time.Sleep(15 * time.Millisecond)
	}
}
