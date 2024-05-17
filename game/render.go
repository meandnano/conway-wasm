package main

import (
	"image/color"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/meandnano/conway-web/game/logic"
)

type Renderer struct {
	game *logic.Game

	boardWidthPx  float64
	boardHeightPx float64

	cellWidthPx  float64
	cellHeightPx float64

	aliveColor color.RGBA
	deadColor  color.RGBA
}

type RendererOpt func(r *Renderer)

func WithAliveColor(c color.RGBA) RendererOpt {
	return func(r *Renderer) {
		r.aliveColor = c
	}
}

func WithDeadColor(c color.RGBA) RendererOpt {
	return func(r *Renderer) {
		r.deadColor = c
	}
}

func NewRenderer(g *logic.Game, canvasW, canvasH float64, opts ...RendererOpt) *Renderer {
	r := &Renderer{
		game: g,

		boardWidthPx:  canvasW,
		boardHeightPx: canvasH,

		cellWidthPx:  canvasW / float64(g.WidthInCells()),
		cellHeightPx: canvasH / float64(g.HeightInCells()),

		deadColor:  color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
		aliveColor: color.RGBA{0, 0, 0, 0xFF},
	}

	for _, o := range opts {
		o(r)
	}

	return r
}

// Render renders the game
func (r *Renderer) Render(gc *draw2dimg.GraphicContext) bool {
	gc.Clear()

	// render the game board
	for xCell := uint(0); xCell < r.game.WidthInCells(); xCell++ {
		r.game.EachColumn(xCell, func(y uint, c *logic.Cell) {
			cx := float64(xCell) * r.cellWidthPx
			cy := float64(y) * r.cellHeightPx

			gc.SetStrokeColor(color.Transparent)
			if c.Alive {
				gc.SetFillColor(r.aliveColor)
			} else {
				gc.SetFillColor(r.deadColor)
			}

			gc.BeginPath()
			gc.MoveTo(cx, cy)
			gc.LineTo(cx+r.cellWidthPx, cy)
			gc.LineTo(cx+r.cellWidthPx, cy+r.cellHeightPx)
			gc.LineTo(cx, cy+r.cellHeightPx)
			gc.Close()
			gc.FillStroke()
		})
	}

	return true
}
