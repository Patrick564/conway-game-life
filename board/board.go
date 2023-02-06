package board

import (
	"fmt"

	color "github.com/fatih/color"
	"github.com/muesli/termenv"
)

const (
	width  = 75 + 2
	height = 40 + 2
)

// ■ \u258 █
var (
	squareFill  = color.WhiteString("██")
	squareEmpty = color.BlackString("██")

	top         = "──"
	left        = "│"
	topLeft     = "╭"
	topRigth    = "╮"
	bottomLeft  = "╰"
	bottomRigth = "╯"
)

type Cell struct {
	Status bool
	Value  int
}

func (c Cell) Live() Cell {
	return Cell{true, 1}
}

func (c Cell) Dead() Cell {
	return Cell{false, 0}
}

type Board struct {
	Data [][]Cell
}

func (b *Board) CheckNeighboursAlive(row, col int) int {
	neighboursAlive := 0

	neighboursAlive += b.Data[row-1][col-1].Value
	neighboursAlive += b.Data[row-1][col].Value
	neighboursAlive += b.Data[row-1][col+1].Value

	neighboursAlive += b.Data[row][col-1].Value
	neighboursAlive += b.Data[row][col+1].Value

	neighboursAlive += b.Data[row+1][col-1].Value
	neighboursAlive += b.Data[row+1][col].Value
	neighboursAlive += b.Data[row+1][col+1].Value

	return neighboursAlive
}

func (b *Board) Draw(output *termenv.Output) {
	for rowIdx, row := range b.Data {
		for colIdx, col := range row {
			if rowIdx == 0 || rowIdx == len(b.Data)-1 {
				if rowIdx == 0 && colIdx == 0 {
					fmt.Fprint(output, topLeft)
					continue
				}

				if rowIdx == 0 && colIdx == len(row)-1 {
					fmt.Fprint(output, topRigth)
					continue
				}

				if rowIdx == len(b.Data)-1 && colIdx == 0 {
					fmt.Fprint(output, bottomLeft)
					continue
				}

				if rowIdx == len(b.Data)-1 && colIdx == len(row)-1 {
					fmt.Fprint(output, bottomRigth)
					continue
				}

				fmt.Fprint(output, top)
				continue
			}

			if colIdx == 0 || colIdx == len(row)-1 {
				fmt.Fprint(output, left)
				continue
			}

			if col.Status {
				fmt.Fprint(output, squareFill)
				continue
			}

			fmt.Fprint(output, squareEmpty)
		}
		fmt.Print("\n")
	}
}

func emptyPattern() Board {
	b := make([][]Cell, height)
	for i := range b {
		c := make([]Cell, width)

		for ci := range c {
			c[ci] = Cell{}
		}

		b[i] = c
	}
	return Board{Data: b}
}

func pulsarPattern() Board {
	b := make([][]Cell, height)
	for i := range b {
		c := make([]Cell, width)

		for ci := range c {
			c[ci] = Cell{}
		}

		if i == 14 {
			c[34] = Cell{true, 1}
			c[35] = Cell{true, 1}
			c[36] = Cell{true, 1}
			c[40] = Cell{true, 1}
			c[41] = Cell{true, 1}
			c[42] = Cell{true, 1}
		}
		if i == 16 {
			c[32] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[44] = Cell{true, 1}
		}
		if i == 17 {
			c[32] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[44] = Cell{true, 1}
		}
		if i == 18 {
			c[32] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[44] = Cell{true, 1}
		}
		if i == 19 {
			c[34] = Cell{true, 1}
			c[35] = Cell{true, 1}
			c[36] = Cell{true, 1}
			c[40] = Cell{true, 1}
			c[41] = Cell{true, 1}
			c[42] = Cell{true, 1}
		}

		if i == 21 {
			c[34] = Cell{true, 1}
			c[35] = Cell{true, 1}
			c[36] = Cell{true, 1}
			c[40] = Cell{true, 1}
			c[41] = Cell{true, 1}
			c[42] = Cell{true, 1}
		}
		if i == 22 {
			c[32] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[44] = Cell{true, 1}
		}
		if i == 23 {
			c[32] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[44] = Cell{true, 1}
		}
		if i == 24 {
			c[32] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[44] = Cell{true, 1}
		}
		if i == 26 {
			c[34] = Cell{true, 1}
			c[35] = Cell{true, 1}
			c[36] = Cell{true, 1}
			c[40] = Cell{true, 1}
			c[41] = Cell{true, 1}
			c[42] = Cell{true, 1}
		}

		b[i] = c
	}
	return Board{Data: b}
}

func pentaPattern() Board {
	b := make([][]Cell, height)
	for i := range b {
		c := make([]Cell, width)

		for ci := range c {
			c[ci] = Cell{}
		}

		if i == 15 {
			c[38] = Cell{true, 1}
		}
		if i == 16 {
			c[37] = Cell{true, 1}
			c[38] = Cell{true, 1}
			c[39] = Cell{true, 1}
		}
		if i == 17 {
			c[36] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[38] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[40] = Cell{true, 1}
		}

		if i == 24 {
			c[36] = Cell{true, 1}
			c[37] = Cell{true, 1}
			c[38] = Cell{true, 1}
			c[39] = Cell{true, 1}
			c[40] = Cell{true, 1}
		}
		if i == 25 {
			c[37] = Cell{true, 1}
			c[38] = Cell{true, 1}
			c[39] = Cell{true, 1}
		}
		if i == 26 {
			c[38] = Cell{true, 1}
		}

		b[i] = c
	}
	return Board{Data: b}
}

func New(pattern string) Board {
	var b Board
	switch pattern {
	case "empty":
		b = emptyPattern()
	case "penta":
		b = pentaPattern()
	case "pulsar":
		b = pulsarPattern()
	}
	return b
}
