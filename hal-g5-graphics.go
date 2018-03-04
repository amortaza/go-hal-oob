package hal_g5

import (
	"github.com/amortaza/go-hal"
	"github.com/amortaza/go-g5"
)

type Graphics struct {

}

func (graphics *Graphics) NewCanvas(width, height int) hal.Canvas {

	return g5.NewCanvas(width, height)
}

func (graphics *Graphics) Clear(r, g, b float32) {

	g5.Clear(r, g, b)
}

func (graphics *Graphics) PushView(width, height int) {

	g5.PushView(width, height)
}

func (graphics *Graphics) PopView() {

	g5.PopView()
}

