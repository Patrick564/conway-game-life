package main

import (
	"os"
	"time"

	"github.com/Patrick564/conway-game-life/board"
	"github.com/muesli/termenv"
)

func main() {
	b := board.New("pulsar")
	output := termenv.NewOutput(os.Stdout)
	output.AltScreen()

	for {
		tmpBoard := board.New("empty")

		for rIdx, row := range b.Data {
			if rIdx == 0 || rIdx == len(b.Data)-1 {
				continue
			}

			for cIdx, col := range row {
				if cIdx == 0 || cIdx == len(row)-1 {
					continue
				}

				neighboursAlive := b.CheckNeighboursAlive(rIdx, cIdx)

				if col.Status && (neighboursAlive == 2 || neighboursAlive == 3) {
					tmpBoard.Data[rIdx][cIdx] = board.Cell{Status: true, Value: 1}
					continue
				}

				if !col.Status && neighboursAlive == 3 {
					tmpBoard.Data[rIdx][cIdx] = board.Cell{Status: true, Value: 1}
					continue
				}

				tmpBoard.Data[rIdx][cIdx] = board.Cell{Status: false, Value: 0}
			}
		}

		b = tmpBoard
		b.Draw(output)
		time.Sleep(500 * time.Millisecond)
		output.ClearScreen()
	}
}
