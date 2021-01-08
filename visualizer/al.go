package visualizer

import (
	"time"

	"github.com/gdamore/tcell/v2"
)


type MovingRectangle struct {
	x        int
	y        int
	initialX    int
	initialY    int
	diameter int
	velx     int
	vely     int
}

type Visualizer struct {
	Screen tcell.Screen
	Rect *MovingRectangle
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
		&MovingRectangle{},
		time.NewTicker(1000000 / 1 * time.Microsecond),
	}, nil
}

func NewMovingRectangle(x int, y int, d int, velx int, vely int) *MovingRectangle {
	mr := &MovingRectangle{x, y, x, y, d, velx, vely}
	return mr
}


func Move(b *MovingRectangle) {
	b.x -= 2 * b.velx
	b.y += b.vely
}


func DrawRect(v *Visualizer, b *MovingRectangle) {
	for i := b.x; i < b.x+2*b.diameter; i++ {
		for j := b.y; j < b.y+b.diameter; j++ {
			// combc is in most cases nil
			v.Screen.SetContent(i, j, 'X', nil, tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite))
		}
	}
}
