package hal_g5

import (
	"github.com/amortaza/go-xel"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-g5"
	"fmt"
)

type Hal struct {
}

func New() *Hal {
	return &Hal{}
}

func (hal *Hal) Start(	width, height int,
			title string,
			onAfterGL,
			onLoop,
			onBeforeDelete func(),
			onResize,
			onMouseMove func(int,int),
			onMouseButton func(bl.MouseButton, bl.ButtonAction),
			onKey func(bl.KeyboardKey, bl.ButtonAction, bool, bool, bool)) {

	xel.Init(title, width, height)

	xel.SetCallbacks(
		func() {
			fmt.Println("(+) g5.Init()")

			g5.Init()

			if onAfterGL != nil {
				onAfterGL()
			}
		},

		onLoop,

		func() {
			if onBeforeDelete != nil {
				onBeforeDelete()
			}

			fmt.Println("(-) g5.Uninit()")

			g5.Uninit()
		},

		onResize, onMouseMove, onMouseButton, onKey)

	xel.Loop()

	xel.Uninit()
}

func (hal *Hal) NewCanvas(width, height int) bl.Canvas {
	return g5.NewCanvas(width, height)
}

func (hal *Hal) Clear(r, g, b, a float32) {
	g5.Clear(r, g, b, a)
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
