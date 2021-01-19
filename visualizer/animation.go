package visualizer

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

//Signal is a way to interact with a running animation
type Signal int

const (
	Stop Signal = iota
	Start 
)

//Animation represents an animation that can be run on a terminal window
type Animation interface {
	//Start starts the Animation and returns a Channel where the client can send signals to interact with the Running animation
	Run(scr tcell.Screen, ticker *time.Ticker, signal <-chan Signal) 
}

//AnimationState ...
type AnimationState struct {
	IsRunning bool
}