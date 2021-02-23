package rectangle

import ( 
   "testing"
   "al.go/terminal"
   "al.go/testing/mocks"
)

func TestMove(t *testing.T) {
	originalPoint := terminal.Point{X: 2, Y: 2,}
	rect := Rectangle{TopLeftCorner: originalPoint, Width: 10, Height: 10,}
	velX, velY := 5, 8
	destinationPoint := terminal.Point{X: originalPoint.X + velX, Y: originalPoint.Y + velY,}
	rect.Move(velX, velY)
	if rect.TopLeftCorner != destinationPoint {
		t.Errorf("Unexpected destination point after moving the rectangle: expected %v, got: %v", destinationPoint, rect.TopLeftCorner)
	}

	if rect.Height != 10 {
		t.Errorf("Changing the Rectangle's height is not allowed while moving")
	}

	if rect.Width != 10 {
		t.Errorf("Changing the Rectangle's width is not allowed while moving")
	}
}

func TestHasHitScreenBorder(t *testing.T) {
	scr := &mocks.ScreenMock{
		SizeFunc: func() (int, int) {
			return 50, 50
		},
	}

	w, h := 5, 5

	points := []terminal.Point{ 
								{X: 10, Y: 10}, 
								{X: 0, Y: 10},
								{X: 45, Y:10},
								{X: 10, Y: 0},
								{X: 10, Y: 45},
							  }
	
	expectedResults := []struct{
			any bool
			top bool
			bottom bool
			left bool
			right bool
	}{
		{ any: false, top: false, bottom: false, left: false, right: false },
		{ any: true, top: false, bottom: false, left: true, right: false },
		{ any: true, top: false, bottom: false, left: false, right: true },
		{ any: true, top: true, bottom: false, left: false, right: false },
		{ any: true, top: false, bottom: true, left: false, right: false },
	}

	for i, p := range points {
		rect := Rectangle{TopLeftCorner: p, Width: w, Height: h, }
		expectedResult := expectedResults[i]
		
		any, top, bottom, left, right := rect.HasHitScreenBorder(scr)

		if expectedResult.any == any &&
		   expectedResult.top == top &&
		   expectedResult.bottom == bottom &&
		   expectedResult.left == left &&
		   expectedResult.right == right {
			   continue
		   } else {
			   t.Errorf("Invalid result for HasHitScreenBorder at point: %v \n(any, top, bottom, left, right)\nexpected:\n%v, %v, %v, %v, %v\ngot:\n%v, %v, %v, %v, %v", p,
			   any, top, bottom, left, right, 
			   expectedResult.any, expectedResult.top, expectedResult.bottom, expectedResult.left, expectedResult.right)
		   }
	}
}

func TestDraw(t *testing.T) {
	scr := &mocks.ScreenMock{
		DrawAtFunc: func(p terminal.Point, options terminal.DrawningOptions) {
		},
	}
	originalPoint := terminal.Point{X: 2, Y: 2,}
	rect := Rectangle{TopLeftCorner: originalPoint, Width: 10, Height: 10,}
	//Since a character on the terminal is slightly taller than it is broader, 
	//we must occupy two characters on the width of the rectangle to get better proportions
	area := (rect.Width * 2) * rect.Height

	rect.Draw(scr, terminal.ColorBlue)

	if len(scr.DrawAtCalls()) != area {
		t.Errorf("Calls to method DrawAt() are different than the rectangle Area. expected %v, got %v", area, len(scr.DrawAtCalls()))
	}
}