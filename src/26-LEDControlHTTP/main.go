package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/stianeikeland/go-rpio"
)

var ledPin rpio.Pin

func LEDON(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LED ON")
	fmt.Println(ledPin)
	ledPin.High()
}

func LEDOFF(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LED OFF")
	fmt.Println(ledPin)
	ledPin.Low()
}

func InitGPIO() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get the pin for each of the lights
	ledPin := rpio.Pin(4)
	fmt.Println(ledPin)

	// Set the pins to output mode
	ledPin.Output()
	fmt.Println(ledPin)

	// Clean up on ctrl-c and turn lights out
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ledPin.Low()
		fmt.Println("Exiting...")
		os.Exit(0)
	}()

	// Turn lights off to start.
	ledPin.High()
	fmt.Println("LED Initialized")
}

func main() {

	http.HandleFunc("/ledon", LEDON)
	http.HandleFunc("/ledoff", LEDOFF)

	// if err := rpio.Open(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	defer rpio.Close()

	InitGPIO()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Server Crashed...")
	}

}
