package rectangle

import ( 
   "testing"
   "al.go/terminal"
)

func TestMove(t *testing.T) {
	originalPoint := terminal.Point{X: 2, Y: 2,}
	rect := Rectangle{TopLeftCorner: originalPoint, Width: 10, Height: 10,}
	velX, velY := 5, 8
	destinationPoint := terminal.Point{X: originalPoint.X + velX, Y: originalPoint.Y + velY,}
	rect.Move(velX, velY)
	if rect.TopLeftCorner != destinationPoint {
		t.Errorf("Unexpected destination point after moving rectangle: expected %v, got: %v", destinationPoint, rect.TopLeftCorner)
	}

	if rect.Height != 10 {
		t.Errorf("Changing the Rectangle's height is not allowed while moving")
	}

	if rect.Width != 10 {
		t.Errorf("Changing the Rectangle's width is not allowed while moving")
	}
}

func TestHasHitScreenBorder(t *testing.T) {
	


}