package main

import (
	"al.go/terminal"
	"al.go/terminal/tcell"
	"al.go/visualizer"
	"al.go/visualizer/misc"
)

func main() {
	tc := tcell.New()
	v := visualizer.New(tc)

	mv := misc.NewMovingRectangleAnimation()
	keyEvents := make(chan terminal.KeyboardEvent)

	go tc.WaitForEvent(keyEvents)

	v.Visualize(mv, keyEvents)
}