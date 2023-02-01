package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Cell struct {
	Status bool
}

type Board [][]Cell

func (b *Board) CheckNeighboursAlive(row, col int) int {
	neighboursAlive := 0

	row -1 row row + 1

	return 0
}

func NewBoard() Board {
	b := make(Board, 7)
	for i := range b {
		b[i] = make([]Cell, 7)
	}
	return b
}

func main() {
	b := NewBoard()
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	b[3][2].Status = true
	b[3][3].Status = true
	b[3][4].Status = true

	for {
		for rIdx, row := range b {
			for cIdx, col := range row {
				neighboursAlive := b.CheckNeighboursAlive(rIdx, cIdx)

				if col.Status && neighboursAlive > 3 {
					col.Status = !col.Status
				}
			}
		}

		time.Sleep(2 * time.Second)
		for _, row := range b {
			fmt.Printf("%+v\n", row)
		}
		cmd.Run()
	}
}
