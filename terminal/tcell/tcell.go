/*Package tcell wraps "github.com/gdamore/tcell/v2" package for drawning on the terminal
This is an implementation of the terminal.Screen and terminal.KeyboardListener interfaces
*/
package tcell

import (
	"al.go/terminal"
	"github.com/gdamore/tcell/v2"
)

//Terminal is the implementation of the terminal interfaces using the tcell library
type Terminal struct {
	scr tcell.Screen
}

//New creates a tcell terminal
func New() *Terminal {
	s, _ := tcell.NewScreen()

	return &Terminal{s}
}

//Init initializes the screen for use
func (t *Terminal) Init() error {
	if err := t.scr.Init(); err != nil {
		return err
	}

	return nil
}

//WaitForEvent is the implementation of the terminal.KeyboardListener interface
func (t *Terminal) WaitForEvent(events chan<- terminal.KeyboardEvent) {
	for {
		ev := t.scr.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			events <- terminal.KeyboardEvent{
				KeyPressed: terminal.Key(tcell.KeyNames[ev.Key()]),
			}
		}
	}
}

//Size returns the Size of the tcell screen
func (t *Terminal) Size() (int, int) {
	return t.scr.Size()
}

//Clear the tcell screen
func (t *Terminal) Clear() {
	t.scr.Clear()
}

//Show the content on the screen
func (t *Terminal) Show() {
	t.scr.Show()
}

//DrawAt uses the tcell.Screen.SetContent() to draw at a specific point
func (t *Terminal) DrawAt(p terminal.Point, s terminal.DrawningOptions) {
	t.scr.SetContent(p.X, p.Y, s.Primary, nil, tcell.StyleDefault.Background((tcell.Color(s.BackgroudColor))).Foreground(tcell.Color(s.ForegroundColor)))
}
