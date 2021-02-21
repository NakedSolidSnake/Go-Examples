package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get the pin for each of the lights
	LEDPin := rpio.Pin(4)

	// Set the pins to output mode
	LEDPin.Output()

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		LEDPin.Low()
		fmt.Println("Exiting...")
		os.Exit(0)
	}()

	defer rpio.Close()

	// Turn lights off to start.
	LEDPin.Low()

	// A while true loop.
	for {
		LEDPin.High()
		time.Sleep(time.Second * 1)

		LEDPin.Low()
		time.Sleep(time.Second * 1)
	}
}
