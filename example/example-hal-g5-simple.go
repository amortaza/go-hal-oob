package main

import (
	"runtime"
	"github.com/amortaza/go-hal-g5"
)

var hal *hal_g5.Hal

var first = true

func onLoop() {

	if first {
		first = false
	} else {
		return
	}
	//hal.GetGraphics().Clear(1,0,0)

	w, h := hal.GetWindowDim()
	g_graphics := hal.GetGraphics()

	g_graphics.PushView(w, h)

	g_graphics.Clear(.3, .3, .3)

	//canvas := renderCanvas(Root_Node)

	//canvas.Paint(false, 0, 0, four_ones_float32)

	g_graphics.PopView()
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
