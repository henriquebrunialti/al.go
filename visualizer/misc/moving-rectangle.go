package misc

import (
	"time"

	"github.com/gdamore/tcell/v2"

	"al.go/visualizer"
	"al.go/visualizer/objects"
)

//MovingRectangle is a Animation of a Rectangle Moving Around the screen
type MovingRectangle struct {
	rect *objects.Rectangle
	state visualizer.AnimationState
}

func newMovingRectangle(x int, y int, w, h int) *objects.Rectangle {
	mr := objects.New(objects.Point{X: x, Y: y}, w, h)
	return mr
}

//NewMovingRectangleAnimation creates the moving rectangle animation
func NewMovingRectangleAnimation() *MovingRectangle {
	return &MovingRectangle{
		state: visualizer.AnimationState { IsRunning: false },
	}
}

//Run the animation
func (mv *MovingRectangle) Run(scr tcell.Screen, ticker *time.Ticker, signal <-chan visualizer.Signal)  {
	w, _ := scr.Size()
	if w/2%2 == 1 {
		w += 2
	}
	mv.state.IsRunning = true
	mv.rect = newMovingRectangle((w-1)/2, 0, 5,5)
	for {
		select {
		case <-ticker.C:
			if mv.state.IsRunning {
				scr.Clear()
				mv.rect.MoveDown(2)
				mv.rect.Draw(scr)
				scr.Show()
			}
		case s := <-signal: 
			mv.handleSignal(s)
		}	
	}
}

func (mv *MovingRectangle) handleSignal(s visualizer.Signal){

	switch s {
	case visualizer.Stop:
		mv.state.IsRunning = false
	case visualizer.Start:
		mv.state.IsRunning = true
	}
}