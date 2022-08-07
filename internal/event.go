package internal

type Event int

const (
	EventCellUpdated Event = iota + 1
	EventGameWon
	EventDraw
)
