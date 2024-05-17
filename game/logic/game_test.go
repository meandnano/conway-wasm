package logic_test

import (
	"testing"

	"github.com/meandnano/conway-web/game/logic"
)

func TestNeigbours(t *testing.T) {
	t.Run("empty-board", func(t *testing.T) {
		g := logic.NewGame(0, 0)
		g.EachNeighbour(1, 1, func(x, y int32, c *logic.Cell) {
			t.Fatalf("neighbour found on empty board at [%d,%d]", x, y)
		})
	})

	t.Run("no-neighbours", func(t *testing.T) {
		g := logic.NewGame(1, 1)
		g.EachNeighbour(0, 0, func(x, y int32, c *logic.Cell) {
			t.Fatalf("neighbour found on single-cell board at [%d,%d]", x, y)
		})
	})

	t.Run("single-neighbour", func(t *testing.T) {
		g := logic.NewGame(1, 2) // 2 cells
		g.EachNeighbour(0, 0, func(x, y int32, c *logic.Cell) {
			if x != 0 || y != 1 {
				t.Fatalf("wrong neighbour found at [%d,%d], expected only [0,1]", x, y)
			}
		})
	})

	t.Run("eight-neighbous", func(t *testing.T) {
		// o o o
		// o x o
		// o o o

		g := logic.NewGame(3, 3) // 2 cells
		neighbourCount := 0
		g.EachNeighbour(1, 1, func(x, y int32, c *logic.Cell) {
			neighbourCount++
			switch neighbourCount {
			case 1:
				assertLocation(t, 0, 0, x, y)
			case 2:
				assertLocation(t, 0, 1, x, y)
			case 3:
				assertLocation(t, 0, 2, x, y)
			case 4:
				assertLocation(t, 1, 0, x, y)
			case 5:
				assertLocation(t, 1, 2, x, y)
			case 6:
				assertLocation(t, 2, 0, x, y)
			case 7:
				assertLocation(t, 2, 1, x, y)
			case 8:
				assertLocation(t, 2, 2, x, y)
			default:
				t.Fatal("expected 8 neignours, got more")
			}
		})
	})

	t.Run("edge-neighbours", func(t *testing.T) {
		// o o o
		// o o x
		// o o o
		g := logic.NewGame(3, 3) // 2 cells
		neighbourCount := 0
		g.EachNeighbour(1, 2, func(x, y int32, c *logic.Cell) {
			neighbourCount++
			switch neighbourCount {
			case 1:
				assertLocation(t, 0, 1, x, y)
			case 2:
				assertLocation(t, 0, 2, x, y)
			case 3:
				assertLocation(t, 1, 1, x, y)
			case 4:
				assertLocation(t, 2, 1, x, y)
			case 5:
				assertLocation(t, 2, 2, x, y)
			default:
				t.Fatal("expected 5 neignours, got more")
			}
		})
	})

	t.Run("corner-neighbours", func(t *testing.T) {
		// o o o
		// o o o
		// o o x
		g := logic.NewGame(3, 3) // 2 cells
		neighbourCount := 0
		g.EachNeighbour(2, 2, func(x, y int32, c *logic.Cell) {
			neighbourCount++
			switch neighbourCount {
			case 1:
				assertLocation(t, 1, 1, x, y)
			case 2:
				assertLocation(t, 1, 2, x, y)
			case 3:
				assertLocation(t, 2, 1, x, y)
			default:
				t.Fatal("expected 3 neignours, got more")
			}
		})
	})
}

func TestTraverse(t *testing.T) {
	t.Run("empty-board", func(t *testing.T) {
		g := logic.NewGame(0, 0)
		g.Traverse(func(x, y int32, c *logic.Cell) {
			t.Fatalf("cell found on empty board at [%d,%d]", x, y)
		})
	})

	t.Run("single-cell", func(t *testing.T) {
		g := logic.NewGame(1, 1)
		g.Traverse(func(x, y int32, c *logic.Cell) {
			assertLocation(t, 0, 0, x, y)
		})
	})

	t.Run("proper-board", func(t *testing.T) {
		g := logic.NewGame(2, 2) // 2 cells
		cellCount := 0
		g.EachNeighbour(1, 1, func(x, y int32, c *logic.Cell) {
			cellCount++
			switch cellCount {
			case 1:
				assertLocation(t, 0, 0, x, y)
			case 2:
				assertLocation(t, 0, 1, x, y)
			case 3:
				assertLocation(t, 1, 0, x, y)
			case 4:
				assertLocation(t, 1, 1, x, y)
			default:
				t.Fatal("expected 4 cells, got more")
			}
		})
	})

	t.Run("uneven-board", func(t *testing.T) {
		g := logic.NewGame(2, 3) // 2 cells
		cellCount := 0
		g.Traverse(func(x, y int32, c *logic.Cell) {
			cellCount++
			switch cellCount {
			case 1:
				assertLocation(t, 0, 0, x, y)
			case 2:
				assertLocation(t, 0, 1, x, y)
			case 3:
				assertLocation(t, 0, 2, x, y)
			case 4:
				assertLocation(t, 1, 0, x, y)
			case 5:
				assertLocation(t, 1, 1, x, y)
			case 6:
				assertLocation(t, 1, 2, x, y)
			default:
				t.Fatal("expected 6 cells, got more")
			}
		})
	})
}

// TestsGameRules
// https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life#Rules
func TestCellReact(t *testing.T) {
	tests := []struct {
		name       string
		alive      bool
		neighbours []uint8
		wantAlive  bool
		wantResult logic.CellResult
	}{
		{
			name:       "underpopulation",
			alive:      true,
			neighbours: []uint8{0, 1},
			wantAlive:  false,
			wantResult: logic.DIED,
		},
		{
			name:       "keep-living",
			alive:      true,
			neighbours: []uint8{2, 3},
			wantAlive:  true,
			wantResult: logic.SAME,
		},
		{
			name:       "overpopulation",
			alive:      true,
			neighbours: []uint8{4, 5, 6, 7, 8, 9}, // 9 is unrealistic, but just in case
			wantAlive:  false,
			wantResult: logic.DIED,
		},
		{
			name:       "birth",
			alive:      false,
			neighbours: []uint8{3},
			wantAlive:  true,
			wantResult: logic.BORN,
		},
		{
			name:       "keeps-dead",
			alive:      false,
			neighbours: []uint8{1, 2, 4, 5, 6, 7, 8, 9}, // 9 is unrealistic, but just in case
			wantAlive:  false,
			wantResult: logic.SAME,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, n := range tt.neighbours {
				cell := &logic.Cell{Alive: tt.alive}
				gotResult := cell.React(n)
				if gotResult != tt.wantResult {
					t.Errorf("unexpected cell result %d, wanted %d (was alive=%t with %d neighbours)", gotResult, tt.wantResult, tt.alive, n)
				}

				if cell.Alive != tt.wantAlive {
					t.Errorf("unexpected cell state alive=%t, wanted alive=%t (was alive=%t with %d neighbours)", cell.Alive, tt.wantAlive, tt.alive, n)
				}
			}
		})
	}
}

func assertLocation(t *testing.T, expX, expY, gotX, gotY int32) {
	if expX != gotX || expY != gotY {
		t.Fatalf("expected location [%d,%d], got [%d,%d]", expX, expY, gotX, gotY)
	}
}
