package rectangle

import (

	"al.go/terminal"
)

//Rectangle is an object that can be drawn on the screen
type Rectangle struct {
	Width         int
	Height        int
	TopLeftCorner terminal.Point
	widthOffset   int
}

//Draw the Rectangle on screen
func (rect *Rectangle) Draw(scr terminal.Screen, color terminal.Color) {
	style := terminal.DrawningOptions{
		Primary:         ' ',
		BackgroudColor:  color,
		ForegroundColor: terminal.ColorWhite,
	}

	rect.DrawWithOptions(scr, style)
}

//DrawWithOptions draws the rectangle with a set of drawning options
func (rect *Rectangle) DrawWithOptions(scr terminal.Screen, options terminal.DrawningOptions) {
	//Traverse the width of the Rectangle starting at the top left corner point
	for i := rect.TopLeftCorner.X; i < rect.TopLeftCorner.X+rect.Width*2 + rect.widthOffset; i++ {
		//Traverse the height of the Rectangle starting at top left corner point
		for j := rect.TopLeftCorner.Y; j < rect.TopLeftCorner.Y+rect.Height; j++ {
			//Fill the rectangle with whitespaces
			p := terminal.Point{X: i, Y: j}
			scr.DrawAt(p, options)
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
	bottom = rect.TopLeftCorner.Y+rect.Height >= scrHeight
	left = rect.TopLeftCorner.X <= 1
	right = rect.TopLeftCorner.X+(rect.Width*2) >= scrWidth
	return (top || bottom || left || right), top, bottom, left, right
}

//DrawWithText draws only the rectangle border, and writes the rune in the middle.
func (rect *Rectangle) DrawWithText(scr terminal.Screen, options terminal.DrawningOptions, value rune) {
	//This is necessary because as we draw the rectangle on the Width * 2, we will always get an even number, and the rune will never be centrelized
	//By offsetting the width + 1 we can always have an Odd Number. 
	rect.widthOffset = 1
	rect.DrawWithOptions(scr, options)
	middle := terminal.Point { X: rect.TopLeftCorner.X+((rect.Width*2) + rect.widthOffset) / 2, Y: rect.TopLeftCorner.Y+ rect.Height / 2, }
	opt := terminal.DrawningOptions { Primary: value, BackgroudColor: options.BackgroudColor, ForegroundColor: options.ForegroundColor }
	scr.DrawAt(middle, opt)
}

//DrawBorders draws only the borders of the rectangle
func (rect *Rectangle) DrawBorders(scr terminal.Screen, options terminal.DrawningOptions) {
	point := terminal.Point{X: rect.TopLeftCorner.X, Y: rect.TopLeftCorner.Y,}
	for point.X < rect.TopLeftCorner.X+rect.Width*2 + rect.widthOffset {
		scr.DrawAt(point, options)
		point.X++
	}
	point.X--
	point.Y++
	for point.Y < rect.TopLeftCorner.Y+rect.Height {
		scr.DrawAt(point, options)
		point.Y++
	}
	point.Y--
	point.X--
	for point.X > rect.TopLeftCorner.X {
		scr.DrawAt(point, options)
		point.X--
	}
	for point.Y > rect.TopLeftCorner.Y - 1 {
		scr.DrawAt(point, options)
		point.Y--
	}
}