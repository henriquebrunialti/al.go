package visualizer

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"al.go/visualizer/objects"
)


type Visualizer struct {
	Screen tcell.Screen
	Rect *objects.Rectangle
	Ticker *time.Ticker
}

func New() (*Visualizer, error) {
	s, err := tcell.NewScreen()

	if err != nil {
		return nil, err
	}

	if err = s.Init(); err != nil {
		return nil, err
	}
	
	return &Visualizer{
		s,
		&objects.Rectangle{},
		time.NewTicker(1000000 / 1 * time.Microsecond),
	}, nil
}

func NewMovingRectangle(x int, y int, w, h int) *objects.Rectangle {
	mr := objects.New(objects.Point{x, y}, w, h)
	return mr
}

//Visualize an Animation
func Visualize(animation Animation, v  *Visualizer, quit chan bool) {
	w, _ := v.Screen.Size()
	if w/2%2 == 1 {
		w += 2
	}
	v.Rect = NewMovingRectangle((w-1)/2, 0, 5,5)
	exit := false
	for !exit {
		select {
		case <-v.Ticker.C:
			 v.Screen.Clear()
			 v.Rect.MoveDown(-2)
			 v.Rect.Draw(v.Screen)
			 v.Screen.Show()
		case <-quit:
			exit = true
			break
		}
	}
}