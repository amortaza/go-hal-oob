package hal_g5

import (
	"github.com/amortaza/go-xel"
	"github.com/amortaza/go-g5"
	"github.com/amortaza/go-bellina"
)

type Hal struct {
}

func New() *Hal {

	return &Hal{}
}

func (hal *Hal) Start(
		width, height int,
		title string,
		onAfterGL,
		onLoop,
		onBeforeDeleteWindow func(),
		onResize,
		onMouseMove func(int,int),
		onMouseButton func(xel.MouseButton, xel.ButtonAction),
		onKey func(xel.KeyboardKey, xel.ButtonAction, bool, bool, bool)) {

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

		onResize, onMouseMove, onMouseButton, onKey)

	xel.Loop(title)
}

func (hal *Hal) NewCanvas(width, height int) bl.Canvas {

	return g5.NewCanvas(width, height)
}

func (hal *Hal) Clear(r, g, b float32) {

	g5.Clear(r, g, b)
}

func (hal *Hal) PushView(width, height int) {

	g5.PushView(width, height)
}

func (hal *Hal) PopView() {

	g5.PopView()
}

func (hal *Hal) GetGraphics() bl.Graphics {

	return hal
}

func (hal *Hal) GetWindowDim()(width, height int) {

	return xel.WinWidth, xel.WinHeight
}

func (hal *Hal) GetMousePos()(x,y int) {

	return xel.MouseX, xel.MouseY
}

func (hal *Hal) SetMouseCursor(cursor bl.MouseCursor) {

	xel.SetMouseCursor(cursor)
}
