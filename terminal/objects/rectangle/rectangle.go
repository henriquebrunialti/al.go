package rectangle

import "al.go/terminal"

//Rectangle is an object that can be drawn on the screen
type Rectangle struct {
	Width        int
	Height        int
	TopLeftCorner terminal.Point
}

//Draw the Rectangle on screen
func (rect *Rectangle) Draw(scr terminal.Screen, color terminal.Color) {
	//Traverse the width of the Rectangle starting at the top left corner point
	for i := rect.TopLeftCorner.X; i < rect.TopLeftCorner.X + rect.Width * 2; i++ {
		//Traverse the height of the Rectangle starting at top left corner point
		for j := rect.TopLeftCorner.Y; j < rect.TopLeftCorner.Y + rect.Height; j++ {
			//Fill the rectangle with whitespaces
			p := terminal.Point{ X: i, Y: j}
			style := terminal.DrawningOptions{ 
				Primary: ' ',
				BackgroudColor: color,
				ForegroundColor: terminal.ColorWhite,
			}
			scr.DrawAt(p, style)
		}
	}
}

//Move the rectangle
func (rect *Rectangle) Move(velocityX, velocityY int) {
	rect.TopLeftCorner.X += velocityX
	rect.TopLeftCorner.Y += velocityY
}


//HasHitScreenBorder returns true if any side of the rectangle touchs the screen border
func (rect *Rectangle) HasHitScreenBorder(scr terminal.Screen) (any, top, bottom, left, right bool) {
	scrWidth, scrHeight := scr.Size()
	top = rect.TopLeftCorner.Y <= 1
	bottom = rect.TopLeftCorner.Y + rect.Height >= scrHeight
	left = rect.TopLeftCorner.X <= 1
	right = rect.TopLeftCorner.X + (rect.Width * 2) >= scrWidth
	return (top || bottom || left || right), top, bottom, left, right
}