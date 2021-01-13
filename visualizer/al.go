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