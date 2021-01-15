package objects

import "github.com/gdamore/tcell/v2"

//Rectangle is a object that can be drawn on the screen
type Rectangle struct {
	Width        int
	Height        int
	TopLeftCorner Point
}

//Draw the Rectangle on the screen
func (rect *Rectangle) Draw(scr tcell.Screen) {
	//Traverse the width of the Rectangle starting at the top left corner point
	for i := rect.TopLeftCorner.X; i < rect.TopLeftCorner.X + rect.Width * 2; i++ {
		//Traverse the height of the Rectangle starting at top left corner point
		for j := rect.TopLeftCorner.Y; j < rect.TopLeftCorner.Y + rect.Height; j++ {
			//Fill the rectangle with whitespaces
			scr.SetContent(i, j, ' ', nil, tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite))
		}
	}
}

//MoveDown moves the rectangle downwards
func (rect *Rectangle) MoveDown(velocity int) {
	if velocity < 0 {
		velocity *= -1
	}

	rect.TopLeftCorner.Y += velocity
}

//New creates a new Rectangle
func New(position Point, width, heigth int) *Rectangle{
	return &Rectangle{
		Width: width,
		Height: heigth,
		TopLeftCorner: position,
	}
}