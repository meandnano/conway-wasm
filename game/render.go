package main

import (
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/meandnano/conway-web/game/logic"
)

type Renderer struct {
	game *logic.Game

	boardWidthPx  float64
	boardHeightPx float64

	cellRadiusPx  float64
	cellPaddingPx float64

	cellColor color.RGBA
	bgColor   color.RGBA
}

type RendererOpt func(r *Renderer)

func WithCellColor(c color.RGBA) RendererOpt {
	return func(r *Renderer) {
		r.cellColor = c
	}
}

func WithBgColor(c color.RGBA) RendererOpt {
	return func(r *Renderer) {
		r.bgColor = c
	}
}

func WithCellPaddingPx(p float64) RendererOpt {
	return func(r *Renderer) {
		r.cellPaddingPx = p
	}
}

func NewRenderer(g *logic.Game, canvasW, canvasH float64, opts ...RendererOpt) *Renderer {
	radius := math.Min(canvasW/float64(g.WidthInCells()), canvasH/float64(g.HeightInCells()))

	r := &Renderer{
		game: g,

		boardWidthPx:  canvasW,
		boardHeightPx: canvasH,

		cellRadiusPx:  radius,
		cellPaddingPx: 10,

		bgColor:   color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		cellColor: color.RGBA{0, 0, 0, 0xFF},
	}

	for _, o := range opts {
		o(r)
	}

	return r
}

// Render renders the game
func (r *Renderer) Render(gc *draw2dimg.GraphicContext) bool {
	gc.Clear()
	gc.SetFillColor(r.bgColor)
	draw2dkit.Rectangle(gc, 0, 0, r.boardWidthPx, r.boardHeightPx)
	gc.FillStroke()

	// render the game board
	for xCell := uint(0); xCell < r.game.WidthInCells(); xCell++ {
		r.game.EachColumn(xCell, func(y uint, c *logic.Cell) {
			if !c.Alive {
				return
			}

			cx := float64(xCell)*r.cellRadiusPx + r.cellRadiusPx/2
			cy := float64(y)*r.cellRadiusPx + r.cellRadiusPx/2

			gc.SetFillColor(r.cellColor)
			gc.BeginPath()
			draw2dkit.Circle(gc, cx, cy, r.cellRadiusPx-r.cellPaddingPx/2)
			// gc.MoveTo(cx, cy)
			// gc.LineTo(cx+r.cellWidthPx, cy)
			// gc.LineTo(cx+r.cellWidthPx, cy+r.cellHeightPx)
			// gc.LineTo(cx, cy+r.cellHeightPx)
			// gc.Close()
			gc.FillStroke()
		})
	}

	return true
}
