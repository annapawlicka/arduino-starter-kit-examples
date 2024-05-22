# arduino-fun
Implementing examples from the Arduino Uno Starter Kit in Go, or rather [TinyGo](https://tinygo.org).

## Setup

On MacoOS:
```avrasm
brew tap tinygo-org/tools
brew install tinygo
```

Arduino Uno is AVR based (Harvard architecture) microcontroller, so we need to install AVRDUDE. It's an AVR Downloader/Uploader.
```avrasm
brew tap osx-cross/avr
brew install avr-gcc
brew install avrdude
```

## Examples

### Blinky

Turns the on-board LED on for one second, then off for one second, repeatedly.

### Spaceship Interface

Blinks red LEDs alternatively when button is pressed. Otherwise, green LED is on.

![IMG_5368](https://github.com/annapawlicka/arduino-fun/assets/2522010/ef5abf3b-6fb3-49cf-a762-8c8aa713d569)

### Love-o-meter

Uses analog temperature sensor to measure temperature of your finger. LEDs stay off if your temperature is below 20C.
One red LED lights up if your temperature is higher than 20C but less than 23C. All 3 LEDs light up if your temperature is
higher than 23C.

![IMG_5382](https://github.com/annapawlicka/arduino-fun/assets/2522010/4704eeee-b727-4b2e-9f2a-2b18601abfe8)

### Color mixing lamp

Uses 3 photoresistors assigned to Red, Blue and Green colors, and one RGB led.
It captures the amount of light incident to the sensors and transforms it to an analogical signal.
Then, this signal is sent to the RGB led through digital outputs.
RGB led transforms this signal into a colored light.

![RGB led](https://github.com/annapawlicka/arduino-fun/assets/2522010/ac1a3b59-8c48-42e4-99a7-652eea65422c)
