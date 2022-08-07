package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type button struct {
	w     *widget.Button
	c     *fyne.Container
	b     *board
	value value
	x, y  int
}

func (b *button) click() {
	t := value(b.w.Text)

	if t != empty {
		return
	}

	b.value = turn
	b.w.SetText(string(turn))

	if *debug {
		fmt.Println("\n" + b.b.String())
	}

	checkWinner(b.b)
	changeTurn()
}
