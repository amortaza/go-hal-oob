package hal_g5

import (
	"github.com/amortaza/go-hal"
	"github.com/amortaza/go-xel"
	"github.com/amortaza/go-g5"
)

type Hal struct {

	graphics *Graphics
}

func NewHal() *Hal {

	return &Hal{ graphics: &Graphics{} }
}

func (hal *Hal) Start(
		title string,

		width, height int,

		onAfterGL,
		onLoop,
		onBeforeDeleteWindow func(),

		onWindowResize,

		onMouseMove func(int,int),
		onMouseButton func(hal.MouseButton, hal.ButtonAction),

		onKey func(hal.KeyboardKey, hal.ButtonAction, bool, bool, bool)) {

	xel.Init(width, height)

	xel.SetCallbacks(

		func() {

			g5.Init()

			if onAfterGL != nil {
				onAfterGL()
			}
		},

		onLoop,

		func() {

			if onBeforeDeleteWindow != nil {
				onBeforeDeleteWindow()
			}

			g5.Uninit()
		},

		onWindowResize,

		onMouseMove,
		onMouseButton,
		onKey)

	xel.Loop(title)
}
func (hal *Hal) GetGraphics() hal.Graphics {

	return hal.graphics
}

func (hal *Hal) GetWindowDim()(width, height int) {

	return xel.WinWidth, xel.WinHeight
}

func (hal *Hal) GetMousePos()(x,y int) {

	return xel.MouseX, xel.MouseY
}

func (hal *Hal) SetMouseCursor(cursor hal.MouseCursor) {

	xel.SetMouseCursor(cursor)
}
