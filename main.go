package main

import (
	"al.go/terminal"
	"al.go/terminal/tcell"
	"al.go/visualizer"
	"al.go/visualizer/misc"
)

func main() {
	v, _ := visualizer.New()
	mv := misc.NewMovingRectangleAnimation()
	k := tcell.New(v.Screen)
	keyEvents := make(chan terminal.KeyboardEvent)

	go k.WaitForEvent(keyEvents)

	v.Visualize(mv, keyEvents)
}