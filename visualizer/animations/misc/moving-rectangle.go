package misc

import (
	"time"
	"al.go/terminal"
	"al.go/terminal/objects/rectangle"
	"al.go/visualizer"
)


//MovingRectangle is a Animation of a Rectangle Moving Around the screen
type MovingRectangle struct {
	rect  *rectangle.Rectangle
	state visualizer.AnimationState
}

func newRectangle(x int, y int, w, h int) *rectangle.Rectangle {
	mr := &rectangle.Rectangle{TopLeftCorner: terminal.Point{X: x, Y: y}, Width: w, Height: h,}
	return mr
}

//NewMovingRectangleAnimation creates the moving rectangle animation
func NewMovingRectangleAnimation() *MovingRectangle {
	return &MovingRectangle{
		state: visualizer.AnimationState{IsRunning: false},
	}
}

//Run the animation
func (mv *MovingRectangle) Run(scr terminal.Screen, ticker *time.Ticker, signal <-chan visualizer.Signal) {
	colors := []terminal.Color{ terminal.ColorBlue, terminal.ColorGreen, terminal.ColorYellow, terminal.ColorAqua, terminal.ColorGray }
	currentColorIdx := 0
	w, _ := scr.Size()
	if w/2%2 == 1 {
		w += 2
	}
	velX, velY := 1, 1
	mv.state.IsRunning = true
	mv.rect = newRectangle((w-1)/2, 3, 2, 2)
	for {
		select {
		case <-ticker.C:
			if mv.state.IsRunning {
				scr.Clear()
				velX, velY, currentColorIdx = mv.handleCollision(scr, velX, velY, currentColorIdx, len(colors))
				mv.rect.Move(velX, velY)
				mv.rect.Draw(scr, colors[currentColorIdx])
				scr.Show()
			}
		case s := <-signal:
			mv.handleSignal(s)
		}
	}
}

func (mv *MovingRectangle) handleCollision(scr terminal.Screen, velX, velY, currentColorIdx, lenColors int) (int, int, int) {
	any, top, bottom, left, right := mv.rect.HasHitScreenBorder(scr)
	if any {
		currentColorIdx++
		if currentColorIdx >= lenColors {
			currentColorIdx = currentColorIdx % lenColors
		}
	}
	if top || bottom {
		velY *= -1
	}
	if left || right {
		velX *= -1
	}
	return velX, velY, currentColorIdx	
}

//CurrentState ...
func (mv *MovingRectangle) CurrentState() visualizer.AnimationState {
	return mv.state
}

func (mv *MovingRectangle) handleSignal(s visualizer.Signal) {
	switch s {
	case visualizer.Stop:
		mv.state.IsRunning = false
	case visualizer.Start:
		mv.state.IsRunning = true
	}
}
