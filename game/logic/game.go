package logic

import (
	"fmt"
	"math/rand"
)

type CellResult int

const (
	SAME CellResult = iota + 1
	DIED
	BORN
)

type Cell struct {
	Alive bool
}

// React implements game rules
// https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules
func (c *Cell) React(neighbours uint8) CellResult {
	if c.Alive {
		if neighbours < 2 || neighbours > 3 {
			c.Alive = false
			return DIED
		}
	} else if neighbours == 3 {
		c.Alive = true
		return BORN
	}

	return SAME
}

type Game struct {
	board  [][]*Cell
	width  uint
	height uint
}

func NewGame(w, h uint32) *Game {
	board := make([][]*Cell, int(w))
	for x := uint32(0); x < w; x++ {
		board[x] = make([]*Cell, int(h))
		for y := uint32(0); y < h; y++ {
			board[x][y] = &Cell{}
		}
	}

	return &Game{
		width:  uint(w),
		height: uint(h),
		board:  board,
	}
}

func (g *Game) PopulateRandom(aliveCount uint) error {
	if aliveCount > uint(g.width*g.height) {
		return fmt.Errorf("too many alive cells wanted: %d, only %d cells are available", aliveCount, g.width*g.height)
	}

	for aliveCount > 0 {
		x := rand.Uint32() % uint32(g.width)
		y := rand.Uint32() % uint32(g.height)
		cell := g.board[x][y]
		if !cell.Alive {
			cell.Alive = true
			aliveCount--
		}
	}

	return nil
}

func (g *Game) Cycle() {
	g.Traverse(func(x, y uint, c *Cell) {
		neighbours := uint8(0)
		g.EachNeighbour(x, y, func(x, y uint, c *Cell) {
			neighbours++
		})

		c.React(neighbours)
	})
}

// traverse calls f on each cell on the board width-first
func (g *Game) Traverse(f func(x, y uint, c *Cell)) {
	for x := uint(0); x < g.width; x++ {
		for y := uint(0); y < g.height; y++ {
			f(x, y, g.board[x][y])
		}
	}
}

// neighbours calls f on each neighbour cell of the cell on [x,y] width-first
func (g *Game) EachNeighbour(x, y uint, f func(x, y uint, c *Cell)) {
	for cx := int(x) - 1; cx <= int(x)+1; cx++ {
		for cy := int(y) - 1; cy <= int(y)+1; cy++ {
			if 0 > cx || uint(cx) >= g.width {
				continue
			}
			if 0 > cy || uint(cy) >= g.height {
				continue
			}
			if uint(cx) == x && uint(cy) == y {
				continue
			}

			f(uint(cx), uint(cy), g.board[cx][cy])
		}
	}
}
