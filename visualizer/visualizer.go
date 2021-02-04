package visualizer

import (
	"log"
	"time"

	"al.go/terminal"
)

//Visualizer is an object that orchestrates an animation by handling things such as Frames per second and keyboard events
//It also blocks the program so it does not quit before the animation ends or the user press "Esc"
type Visualizer struct {
	Scr       terminal.Screen
	Ticker    *time.Ticker
	animation Animation

	exit bool
}

//New creates a new visualizer
func New(scr terminal.Screen) *Visualizer {
	return &Visualizer{
		scr,
		time.NewTicker(time.Second / 8),
		nil,
		false,
	}
}

//Visualize an Animation
func (v *Visualizer) Visualize(animation Animation, keyboard <-chan terminal.KeyboardEvent) {
	err := v.Scr.Init()

	if err != nil {
		log.Fatalf("Fatal error, could not initialize the screen: %v", err)
		return
	}

	v.animation = animation

	s := make(chan Signal)

	go animation.Run(v.Scr, v.Ticker, s)
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
	if evt.KeyPressed == "Enter" {
		if v.animation.CurrentState().IsRunning {
			s <- Stop
		} else {
			s <- Start
		}
	}
}
