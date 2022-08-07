package main

type board [][]*button

func (b board) getWinner() value {
	for i := 0; i < len(b); i++ {
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

func (b board) getColWinner(idx int) value {
	if idx < 0 || idx > len(b) {
		return empty
	}

	var player value

	for i := 0; i < len(b); i++ {
		if i == 0 {
			player = b[i][idx].value
			continue
		}
		if b[i][idx].value != player {
			return empty
		}
	}
	return player
}

func (b board) getRowWinner(idx int) value {
	if idx < 0 || idx > len(b) {
		return empty
	}

	var player value

	for i := 0; i < len(b); i++ {
		if i == 0 {
			player = b[idx][i].value
			continue
		}
		if b[idx][1].value != player {
			return empty
		}
	}
	return player
}

func (b board) getDia1Winner() value {
	var player value

	for idx := 0; idx < len(b); idx++ {
		if idx == 0 {
			player = b[idx][idx].value
			continue
		}
		if b[idx][idx].value != player {
			return empty
		}
	}
	return player
}

func (b board) getDia2Winner() value {
	var player value

	for idx := 0; idx < len(b); idx++ {
		if idx == 0 {
			player = b[len(b)-idx-1][idx].value
			continue
		}
		if b[len(b)-idx-1][idx].value != player {
			return empty
		}
	}
	return player
}

func (b board) String() string {
	var s string
	for _, row := range b {
		for _, cell := range row {
			s += string(cell.value)
		}
		s += "\n"
	}
	return s
}
