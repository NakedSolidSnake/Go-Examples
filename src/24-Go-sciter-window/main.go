package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {

	windowRect := sciter.NewRect(400, 300, 200, 200)

	windowFromSciter, _ := window.New(sciter.SW_TITLEBAR|sciter.SW_MAIN, windowRect)

	windowFromSciter.SetTitle(" Test Window ")

	windowFromSciter.Show()

	windowFromSciter.Run()
}
