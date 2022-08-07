package model

import (
	"fmt"

	common "github.com/ilyakaznacheev/ttt/internal"
)

type value string

const (
	empty   value = " "
	player1 value = "x"
	player2 value = "o"
	draw    value = "-"
)

type Opt struct {
	Debug bool
}

type Board struct {
	b      [][]value
	e      chan<- common.Event
	turn   value
	winner value

	opt Opt
}

func NewBoard(size int, e chan<- common.Event, opt Opt) *Board {
	if size <= 1 {
		panic("Board is to small")
	}

	br := make([][]value, 0, size)

	for i := 0; i < size; i++ {
		row := make([]value, 0, size)
		for j := 0; j < size; j++ {
			row = append(row, empty)
		}
		br = append(br, row)
	}

	return &Board{
		opt:    opt,
		e:      e,
		b:      br,
		turn:   player1,
		winner: empty,
	}
}

func (b *Board) Len() int {
	return len(b.b)
}

func (b *Board) Click(x, y int) func() {
	return func() {
		if b.b[x][y] != empty {
			return
		}

		b.b[x][y] = b.turn

		if b.opt.Debug {
			fmt.Printf("=== cell[%d,%d] -> %s ===\n", x, y, b.turn)
			fmt.Println("=== board ===\n" + b.String())
		}

		b.notify(common.EventCellUpdated)

		if b.checkWinner() {
			b.notify(common.EventGameWon)
			return
		}

		if b.checkGameOver() {
			b.notify(common.EventDraw)
			return
		}
		b.changeTurn()
	}
}

func (b *Board) Value(x, y int) string {
	if x < 0 || x >= b.Len() || y < 0 || y >= b.Len() {
		return string(empty)
	}
	return string(b.b[x][y])
}

func (b *Board) Winner() (string, bool) {
	return string(b.winner), b.winner != empty
}

func (b *Board) checkWinner() bool {
	w := b.getWinner()
	if w == empty {
		return false
	}

	if b.opt.Debug {
		fmt.Printf("=== winner ===\nplayer %q\n", string(w))
	}

	b.winner = w

	return true
}

func (b *Board) getWinner() value {
	for i := 0; i < b.Len(); i++ {
		if p := b.getRowWinner(i); p != empty {
			return p
		}
		if p := b.getColWinner(i); p != empty {
			return p
		}
	}
	if p := b.getDia1Winner(); p != empty {
		return p
	}
	if p := b.getDia2Winner(); p != empty {
		return p
	}
	return empty
}

func (b *Board) getColWinner(idx int) value {
	if idx < 0 || idx > b.Len() {
		return empty
	}

	var player value

	for i := 0; i < b.Len(); i++ {
		if i == 0 {
			player = b.b[i][idx]
			continue
		}
		if b.b[i][idx] != player {
			return empty
		}
	}

	if b.opt.Debug && player != empty {
		fmt.Printf("=== col %d winner ===\n", idx)
	}
	return player
}

func (b *Board) getRowWinner(idx int) value {
	if idx < 0 || idx > b.Len() {
		return empty
	}

	var player value

	for i := 0; i < b.Len(); i++ {
		if i == 0 {
			player = b.b[idx][i]
			continue
		}
		if b.b[idx][i] != player {
			return empty
		}
	}

	if b.opt.Debug && player != empty {
		fmt.Printf("=== row %d winner ===\n", idx)
	}
	return player
}

func (b *Board) getDia1Winner() value {
	var player value

	for idx := 0; idx < b.Len(); idx++ {
		if idx == 0 {
			player = b.b[idx][idx]
			continue
		}
		if b.b[idx][idx] != player {
			return empty
		}
	}

	if b.opt.Debug && player != empty {
		fmt.Println("=== diagonal 1 winner ===")
	}
	return player
}

func (b *Board) getDia2Winner() value {
	var player value

	for idx := 0; idx < b.Len(); idx++ {
		if idx == 0 {
			player = b.b[b.Len()-idx-1][idx]
			continue
		}
		if b.b[b.Len()-idx-1][idx] != player {
			return empty
		}
	}

	if b.opt.Debug && player != empty {
		fmt.Println("=== diagonal 1 winner ===")
	}
	return player
}

func (b *Board) checkGameOver() bool {
	for _, row := range b.b {
		for _, cell := range row {
			if cell == empty {
				return false
			}
		}
	}

	if b.opt.Debug {
		fmt.Println("=== game draw ===")
	}
	return true
}

func (b *Board) changeTurn() {
	switch b.turn {
	case player1:
		b.turn = player2
	case player2:
		b.turn = player1
	default:
		b.turn = player1
	}
}

func (b *Board) notify(e common.Event) {
	go func() { b.e <- e }()
}

func (b *Board) String() string {
	var s string
	for _, row := range b.b {
		for _, cell := range row {
			s += string(cell)
		}
		s += "\n"
	}
	return s
}
