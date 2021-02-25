package misc

import (
	"testing"
	"al.go/visualizer"
)


func Test_handleSignal(t *testing.T) {
	mv := NewMovingRectangleAnimation()

	tests := []struct {
		signal visualizer.Signal
		expectedState  visualizer.AnimationState 
	}{
		{visualizer.Pause, visualizer.AnimationState{IsRunning: false}},
		{visualizer.Start, visualizer.AnimationState{IsRunning: true}},
	}

	for _, test := range tests {
		mv.handleSignal(test.signal)

		if mv.state != test.expectedState {
			t.Errorf("Signal generated a unexpected animation state. Expected: %+v, got: %+v", test.expectedState, mv.state)
		}
	}
}