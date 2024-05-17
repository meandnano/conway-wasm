package main

import (
	"fmt"
	"time"

	"github.com/markfarnan/go-canvas/canvas"
	"github.com/meandnano/conway-web/game/logic"
)

var done chan struct{}

var gameUpdateDelay time.Duration = time.Second / 2

func main() {
	cvs, err := canvas.NewCanvas2d(true)
	if err != nil {
		fmt.Printf("error creating canvas: %s\n", err)
		return
	}

	height := float64(cvs.Height())
	width := float64(cvs.Width())

	game := logic.NewGame(100, 120)
	if err := game.PopulateRandom(1000); err != nil {
		fmt.Printf("error populating: %s\n", err)
		return
	}
	r := NewRenderer(game, width, height)

	cvs.Start(30, r.Render)

	go cycleGame(gameUpdateDelay, game)

	<-done
}

func cycleGame(period time.Duration, g *logic.Game) {
	t := time.NewTicker(period)
	for {
		<-t.C
		g.Cycle()
	}
}
