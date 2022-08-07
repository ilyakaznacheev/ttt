package main

import (
	"flag"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type value string

const (
	empty   value = " "
	player1 value = "x"
	player2 value = "o"
)

var turn value
var win fyne.Window

var (
	size  = flag.Int("s", 3, "board size")
	debug = flag.Bool("debug", false, "debug mode")
)

func main() {
	changeTurn()

	flag.Parse()

	a := app.New()
	win = a.NewWindow("TTT")

	grid := initBoard(*size)

	win.SetContent(grid)
	win.ShowAndRun()
}

func initBoard(size int) *fyne.Container {
	if size <= 1 {
		panic("board is to small")
	}

	grid := container.New(layout.NewAdaptiveGridLayout(size))

	br := make(board, 0, size)

	for i := 0; i < size; i++ {
		row := make([]*button, 0, size)
		for j := 0; j < size; j++ {

			b := button{
				value: empty,
				c:     grid,
				b:     &br,
				x:     i,
				y:     j,
			}

			b.w = widget.NewButton(string(empty), b.click)

			grid.Add(b.w)

			row = append(row, &b)
		}
		br = append(br, row)
	}

	return grid
}

func changeTurn() {
	switch turn {
	case player1:
		turn = player2
	case player2:
		turn = player1
	default:
		turn = player1
	}
}

func checkWinner(b *board) {
	w := b.getWinner()
	if w == empty {
		return
	}

	if *debug {
		fmt.Printf("player %q wins!\n", string(w))
	}

	t := widget.NewLabel(fmt.Sprintf("player %q wins!", string(w)))
	p := widget.NewModalPopUp(t, win.Canvas())
	p.Show()
}
