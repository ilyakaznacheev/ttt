package view

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	common "github.com/ilyakaznacheev/ttt/internal"
)

type ModelManager interface {
	Value(int, int) string
	Winner() (string, bool)
	Click(int, int) func()
	Len() int
}

type View struct {
	window  fyne.Window
	buttons [][]*widget.Button
	model   ModelManager
	event   <-chan common.Event
}

func New(m ModelManager, e <-chan common.Event) *View {
	l := m.Len()
	a := app.New()
	win := a.NewWindow("TTT")

	grid := container.New(layout.NewAdaptiveGridLayout(l))

	matrix := make([][]*widget.Button, 0, l)
	for i := 0; i < l; i++ {
		row := make([]*widget.Button, 0, l)
		for j := 0; j < l; j++ {
			b := widget.NewButton(" ", m.Click(i, j))
			row = append(row, b)
			grid.Add(b)
		}
		matrix = append(matrix, row)
	}

	win.SetContent(grid)

	return &View{
		window:  win,
		buttons: matrix,
		model:   m,
		event:   e,
	}
}

func (v *View) Run() {
	go v.watchEvent()
	v.window.ShowAndRun()
}

func (v *View) watchEvent() {
	for {
		select {
		case e, ok := <-v.event:
			if !ok {
				return
			}

			switch e {
			case common.EventCellUpdated:
				v.refresh()
			case common.EventGameWon:
				v.announceWinner()
			case common.EventDraw:
				v.announceDraw()
			}
		}
	}
}

func (v *View) refresh() {
	for i, row := range v.buttons {
		for j, button := range row {
			button.SetText(v.model.Value(i, j))
		}
	}
}

func (v *View) announceWinner() {
	w, ok := v.model.Winner()
	if !ok {
		return
	}

	p := v.exitPopup(fmt.Sprintf("player %q wins!", string(w)))
	p.Show()
}

func (v *View) announceDraw() {
	p := v.exitPopup("no one wins!")
	p.Show()
}

func (v *View) exitPopup(text string) *widget.PopUp {
	t := widget.NewLabel(text)
	b := widget.NewButton("OK", func() { os.Exit(0) })
	c := container.New(layout.NewVBoxLayout(), t, layout.NewSpacer(), b)
	p := widget.NewModalPopUp(c, v.window.Canvas())
	return p
}
