/*Package tcell wraps "github.com/gdamore/tcell/v2" package for drawning on the terminal 
by implementing the terminal interface
*/
package tcell

import (
	"github.com/gdamore/tcell/v2"
	"al.go/terminal"
)

//Terminal is the implementation of the terminal interfaces using the tcell library
type Terminal struct {
	scr tcell.Screen
}

//New creates a tcell terminal
func New(scr tcell.Screen) *Terminal {
	return &Terminal{scr}
}

//WaitForEvent is the implementation of the terminal.KeyboardListener interface
func (t Terminal) WaitForEvent(events chan<- terminal.KeyboardEvent) {
	for {
		ev := t.scr.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			events <- terminal.KeyboardEvent {
				KeyPressed: terminal.Key(tcell.KeyNames[ev.Key()]),  
			}
		}
	}
}