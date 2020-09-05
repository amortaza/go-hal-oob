package main

import (
	"github.com/amortaza/go-hal-oob"
	"runtime"
)

var hal *haloob.Hal

var first = true

func onLoop() {

	if first {
		first = false
	} else {
		return
	}

	w, h := hal.GetWindowDim()
	g_graphics := hal.GetGraphics()

	g_graphics.PushView(w, h)

	g_graphics.Clear(.3, .3, .3)

	g_graphics.PopView()
}

func init() {
	runtime.LockOSThread()
}

func main() {

	hal = haloob.NewHal()

	hal.Start(
		"Bellina",
		1200, 100,
		640, 480,
		nil,
		onLoop,
		nil,
		nil,
		nil,
		nil,
		nil)
}
