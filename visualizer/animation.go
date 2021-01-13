package visualizer

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

//Signal is a way to interact with a running animation
type Signal int

const (
	stop Signal = iota
	pause
	start 
)

//Animation represents a animation that can be run on a terminal window or a physical console
type Animation interface {
	//Start starts the Animation and returns a Channel where the client can send signals to interact with the Running animation
	Run(scr *tcell.Screen, ticker *time.Ticker) (chan<- Signal, error)
}