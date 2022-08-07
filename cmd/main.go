package main

import (
	"flag"

	"github.com/ilyakaznacheev/ttt/internal"
	"github.com/ilyakaznacheev/ttt/internal/model"
	"github.com/ilyakaznacheev/ttt/internal/view"
)

var (
	size  = flag.Int("s", 3, "board size")
	debug = flag.Bool("debug", false, "debug mode")
)

func main() {

	flag.Parse()

	event := make(chan internal.Event)
	defer close(event)

	m := model.NewBoard(*size, event, model.Opt{Debug: *debug})
	v := view.New(m, event)

	v.Run()
}
