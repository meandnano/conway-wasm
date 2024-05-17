package main

import (
	"fmt"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/markfarnan/go-canvas/canvas"
)

var done chan struct{}

var cvs *canvas.Canvas2d
var width float64
var height float64

// This specifies how long a delay between calls to 'render'.     To get Frame Rate,   1s / renderDelay
var renderDelay time.Duration = 20 * time.Millisecond

func main() {

	FrameRate := time.Second / renderDelay
	println("Hello Browser FPS:", FrameRate)
	var err error
	cvs, err = canvas.NewCanvas2d(true)
	if err != nil {
		fmt.Printf("error creating canvas: %s\n", err)
		return
	}

	// cvs, _ = canvas.NewCanvas2d(false)
	// cvs.Create(int(js.Global().Get("innerWidth").Float()*0.9), int(js.Global().Get("innerHeight").Float()*0.9)) // Make Canvas 90% of window size.  For testing rendering canvas smaller than full windows

	height = float64(cvs.Height())
	width = float64(cvs.Width())

	cvs.Start(60, Render)

	//go doEvery(renderDelay, Render) // Kick off the Render function as go routine as it never returns
	<-done
}

// Helper function which calls the required func (in this case 'render') every time.Duration,  Call as a go-routine to prevent blocking, as this never returns
func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

// Called from the 'requestAnnimationFrame' function.   It may also be called seperatly from a 'doEvery' function, if the user prefers drawing to be seperate from the annimationFrame callback
func Render(gc *draw2dimg.GraphicContext) bool {
	// if gs.laserX+gs.directionX > width-gs.laserSize || gs.laserX+gs.directionX < gs.laserSize {
	// 	gs.directionX = -gs.directionX
	// }
	// if gs.laserY+gs.directionY > height-gs.laserSize || gs.laserY+gs.directionY < gs.laserSize {
	// 	gs.directionY = -gs.directionY
	// }
	//
	// gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
	// gc.Clear()
	// // move red laser
	// gs.laserX += gs.directionX
	// gs.laserY += gs.directionY
	//
	// // draws red 🔴 laser
	// gc.SetFillColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	// gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0xff, 0xff})
	//
	// gc.BeginPath()
	// //gc.ArcTo(gs.laserX, gs.laserY, gs.laserSize, gs.laserSize, 0, math.Pi*2)
	// draw2dkit.Circle(gc, gs.laserX, gs.laserY, gs.laserSize)
	// gc.FillStroke()
	// gc.Close()

	return true
}
