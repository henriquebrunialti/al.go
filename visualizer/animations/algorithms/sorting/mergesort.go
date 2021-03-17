package sorting

import (
	"context"
	"time"

	"al.go/terminal"
	"al.go/terminal/objects/array"
	"al.go/visualizer"
)

//MergeSort animation
type MergeSort struct {
	array array.Array
	state visualizer.AnimationState
}

//MS creates a new MergeSort animation
func MS() *MergeSort {
	return &MergeSort{}
}

func (ms *MergeSort) Run(ctx context.Context, scr terminal.Screen, ticker *time.Ticker, signal <-chan visualizer.Signal) {
	ms.state = visualizer.AnimationState{IsRunning: true}

	arr := array.Array{
		TopLeftCorner: terminal.Point{2, 1},
	}

	arr.AddItem('1')
	arr.AddItem('2')
	arr.AddItem('3')


	for {
		select {
		case <-ticker.C:
			if ms.state.IsRunning {
				scr.Clear()
				arr.Draw(scr)
				scr.Show()
			}
		case s := <-signal:
			ms.handleSignal(s)
		case <-ctx.Done():
			return
		}
	}
}

func (ms *MergeSort) handleSignal(s visualizer.Signal) {
	switch s {
	case visualizer.Pause:
		ms.state.IsRunning = false
	case visualizer.Start:
		ms.state.IsRunning = true
	}
}

func (ms *MergeSort) CurrentState() visualizer.AnimationState {
	return ms.state
}
