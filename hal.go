package haloob

import (
	"github.com/amortaza/go-xel2"
)

type Hal struct {
}

func New() *Hal {
	return &Hal{}
}

func (hal *Hal) Start(width, height int, title string, onAfterGL, onLoop, onBeforeDelete func(), onResize, onMouseMove func(int,int), onMouseButton func(xel.MouseButton, xel.ButtonAction), onKey func(xel.KeyboardKey, xel.ButtonAction, bool, bool, bool)) {
	xel.Init(title, width, height)

	xel.SetCallbacks(onAfterGL, onLoop, onBeforeDelete, onResize, onMouseMove, onMouseButton, onKey)

	xel.Loop()

	xel.Uninit()
}

func (hal *Hal) GetWindowDim()(width, height int) {
	return xel.WinWidth, xel.WinHeight
}

func (hal *Hal) GetMousePos()(x,y int) {
	return xel.MouseX, xel.MouseY
}
