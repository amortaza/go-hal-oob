package main

import (
	"runtime"
	"github.com/amortaza/go-hal-g5"
)

var hal *hal_g5.Hal

func onLoop() {

	hal.GetGraphics().Clear(1,0,0)
}

func init() {
	runtime.LockOSThread()
}

func main() {

	hal = hal_g5.NewHal()

	hal.Start(
		"Hello, HAL!",
		640, 480,
		nil,
		onLoop,
		nil,
		nil,
		nil,
		nil,
		nil)
}
