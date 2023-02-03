package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	color "github.com/fatih/color"
)

const (
	SquareFill  = "&#9632"
	SquareEmpty = ""
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

func (b *Board) Draw() {
	for _, col := range b.Data {
		for _, row := range col {
			if row.Status {
				// ■ \u258 █
				fmt.Printf("%s ", color.WhiteString("■"))
				continue
			}
			fmt.Printf("%s ", color.BlackString("■"))
		}
		fmt.Print("\n")
	}
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func NewBoard(size int) Board {
	b := make([][]Cell, size+1)
	for i := range b {
		c := make([]Cell, size+1)

		for ci := range c {
			c[ci] = Cell{}
		}

		if i == 4 {
			c[3] = Cell{true, 1}
			c[4] = Cell{true, 1}
			c[5] = Cell{true, 1}
		}

		b[i] = c
	}
	return Board{Data: b}
}

func NewBoardT(size int) Board {
	b := make([][]Cell, size+1)
	for i := range b {
		c := make([]Cell, size+1)

		for ci := range c {
			c[ci] = Cell{}
		}

		b[i] = c
	}
	return Board{Data: b}
}

func main() {
	board := NewBoard(8)
	// cmd := exec.Command("clear")
	// cmd.Stdout = os.Stdout

	// fmt.Println("■ \u2580")

	board.Draw()
	// fmt.Println(color.BlackString("■"))
	// fmt.Println("---------------------------------------------------- &#9632")

	for {
		tmpBoard := NewBoardT(8)

		for rIdx, row := range board.Data {
			if rIdx == 0 || rIdx == 8 {
				continue
			}

			for cIdx, col := range row {
				if cIdx == 0 || cIdx == 8 {
					continue
				}

				neighboursAlive := board.CheckNeighboursAlive(rIdx, cIdx)

				if col.Status && (neighboursAlive == 2 || neighboursAlive == 3) {
					tmpBoard.Data[rIdx][cIdx] = Cell{true, 1}
					continue
				}

				if !col.Status && neighboursAlive == 3 {
					tmpBoard.Data[rIdx][cIdx] = Cell{true, 1}
					continue
				}

				tmpBoard.Data[rIdx][cIdx] = Cell{false, 0}
			}
		}
		board = tmpBoard

		board.Draw()
		time.Sleep(200 * time.Millisecond)
		ClearScreen()
		// fmt.Println("-------------------------------------------")
		// cmd.Run()
	}
}
