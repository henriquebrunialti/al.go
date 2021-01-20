package visualizer

import (
	"time"

	"al.go/terminal"
)

//Signal is a way to interact with a running animation
type Signal int

const (
	Stop Signal = iota
	Start 
)

//Animation represents an animation that can be run on a terminal window
type Animation interface {
	//Run starts the Animation and receives a Channel where the client can send signals to interact with the Running animation
	//The animation will move everytime the ticker ticks
	//The only way to change the current state of an animation is to send a signal over the signal channel
	Run(scr terminal.Screen, ticker *time.Ticker, signal <-chan Signal) 

	//CurrentState is the only way to get the current state from an Animation.
	CurrentState() AnimationState
}

//AnimationState ...
type AnimationState struct {
	IsRunning bool
}