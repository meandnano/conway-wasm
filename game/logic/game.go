package logic

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
	width  int32
	height int32
}

func NewGame(w, h int32) *Game {
	board := make([][]*Cell, int(w))
	for x := 0; x < int(w); x++ {
		board[x] = make([]*Cell, int(h))
		for y := int32(0); y < h; y++ {
			board[x][y] = &Cell{}
		}
	}

	return &Game{
		width:  w,
		height: h,
		board:  board,
	}
}

func (g *Game) PopulateRandom(aliveCount int) {
}

func (g *Game) Cycle() {
	g.Traverse(func(x, y int32, c *Cell) {
		neighbours := 0
		g.EachNeighbour(x, y, func(x, y int32, c *Cell) {
			neighbours++
		})
	})
}

// traverse calls f on each cell on the board width-first
func (g *Game) Traverse(f func(x, y int32, c *Cell)) {
	for x := int32(0); x < g.width; x++ {
		for y := int32(0); y < g.height; y++ {
			f(x, y, g.board[x][y])
		}
	}
}

// neighbours calls f on each neighbour cell of the cell on [x,y] width-first
func (g *Game) EachNeighbour(x, y int32, f func(x, y int32, c *Cell)) {
	for cx := x - 1; cx <= x+1; cx++ {
		for cy := y - 1; cy <= y+1; cy++ {
			if 0 > cx || cx >= g.width {
				continue
			}
			if 0 > cy || cy >= g.height {
				continue
			}
			if cx == x && cy == y {
				continue
			}

			f(cx, cy, g.board[cx][cy])
		}
	}
}
