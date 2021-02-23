package terminal

//Screen represents a 2-d screen where we can draw objects
type Screen interface {
	Init() error
	Size() (width int, height int)
	Clear()
	DrawAt(p Point, options DrawningOptions)
	Show()
}

//KeyboardListener provides functions to handle keyboard events
type KeyboardListener interface {
	//Wait for event will wait for keyboard events and send then over the given channel
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

//DrawningOptions sets the Primary Rune that will be written into a specific point
type DrawningOptions struct {
	Primary rune

	BackgroudColor  Color
	ForegroundColor Color
}

//Key represents a Key on the keyboard e.g "Esc", "Enter"...
//Todo: Provide a list of every supported key
type Key string

//Color Todo review the way we handle color
type Color uint64

const colorValid = 1 << 32

//Supported colors
const (
	ColorBlack = colorValid + iota
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
