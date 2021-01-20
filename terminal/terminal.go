package terminal

//Key represents a Key on the keyboard e.g "Esc", "Enter"...
//Todo <- Provide a list of every supported key
type Key string

//KeyboardListener provides functions to handle keyboard events
type KeyboardListener interface {
	//Wait for event will wait for keyboard events and send then to the given channel
	WaitForEvent(events chan<- KeyboardEvent)
}

//KeyboardEvent ...
type KeyboardEvent struct {
	KeyPressed Key
}

//Point represents a 2-d point in the terminal
type Point struct {
	X int
	Y int
}

//Screen represents a 2-d screen where we can draw objects
type Screen interface {
	Init() error
	Size() (width int, height int)
	Clear() 
	DrawAt(p Point, options StylingOptions)
	Show()
}


type StylingOptions struct {
	Primary rune

	BackgroudColor Color
	ForegroundColor Color 
}

//Todo review the way to handle color
type Color uint64

const ColorValid = 1 << 32

const (
	ColorBlack = ColorValid + iota
	ColorMaroon
	ColorGreen
	ColorOlive
	ColorNavy
	ColorPurple
	ColorTeal
	ColorSilver
	ColorGray
	ColorRed
	ColorLime
	ColorYellow
	ColorBlue
	ColorFuchsia
	ColorAqua
	ColorWhite
)