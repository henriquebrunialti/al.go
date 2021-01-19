package visualizer

import (
	"time"

	"al.go/terminal"
	"github.com/gdamore/tcell/v2"
)


type Visualizer struct {
	Screen tcell.Screen
	Ticker *time.Ticker

	exit bool
}

func New() (*Visualizer, error) {
	s, err := tcell.NewScreen()

	if err != nil {
		return nil, err
	}

	if err = s.Init(); err != nil {
		return nil, err
	}
	
	return &Visualizer{
		s,
		time.NewTicker(1000000 / 1 * time.Microsecond),
		false,
	}, nil
}


//Visualize an Animation
func (v  *Visualizer) Visualize(animation Animation, keyboard <-chan terminal.KeyboardEvent) {
	s := make(chan Signal)
	go animation.Run(v.Screen, v.Ticker, s)
	for !v.exit {
		select {
		case evt := <-keyboard:
			v.handleKeyboard(&evt, s)
		}
	}
}

func (v *Visualizer) handleKeyboard(evt *terminal.KeyboardEvent, s chan<- Signal) {
	if evt.KeyPressed == "Esc" {
		s <- Stop
		v.exit = true
		return
	}
	if evt.KeyPressed == "Space" {
		s <- Stop
	}
	if evt.KeyPressed == "Enter" {
		s <- Start
	}
}