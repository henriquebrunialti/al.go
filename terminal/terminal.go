package terminal

//Key represents a Key on the keyboard e.g "Esc", "Enter", "A", "a" ...
type Key string

//KeyboardListener provides functions to handle keyboard events
type KeyboardListener interface {
	//Wait for event will wait for keyboard events and send then to the given channel
	WaitForEvent(Events chan<- KeyboardEvent)
}

//KeyboardEvent ...
type KeyboardEvent struct {
	KeyPressed Key
}