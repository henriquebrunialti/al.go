package array

import (
	"al.go/terminal"
	"al.go/terminal/objects/rectangle"
)

const itemSize = 3

//Array is a object that can be drawn on the Screen
type Array struct {
	items []ArrayItems
	TopLeftCorner terminal.Point
}

//ArrayItems is items that can be drawn inside an array
type ArrayItems struct {
	Value rune
	IsSelected bool
	Options terminal.DrawningOptions
}

func (a *Array) AddItem(value rune) {
	newItem := ArrayItems{value, false,terminal.DrawningOptions{Primary: ' ', BackgroudColor: terminal.ColorGray, ForegroundColor: terminal.ColorOlive,}}
	a.items = append(a.items, newItem)
}

func (a *Array) Draw(scr terminal.Screen) {
	topLeftCorner := a.TopLeftCorner
	for _, item := range a.items {
		rect := rectangle.Rectangle {
			Width: 3,
			Height: 3,
			TopLeftCorner: topLeftCorner,
		}
		topLeftCorner.X += 4 * itemSize 
		rect.DrawWithText(scr, item.Options, item.Value)
	}
}