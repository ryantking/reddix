package term

import (
	"fmt"

	"github.com/RyanTKing/reddix/pkg/event"
	"github.com/RyanTKing/reddix/pkg/ui"
	termbox "github.com/nsf/termbox-go"
)

// Newreturns a new term instance with termbox iniitialized
func New() (*Term, error) {
	err := termbox.Init()
	if err != nil {
		return nil, err
	}

	termbox.SetInputMode(termbox.InputEsc)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCursor(0, 0)

	w, h := termbox.Size()
	term := Term{
		width:  w,
		height: h,
	}

	return &term, nil
}

// Close closes the termbox section
func (*Term) Close() {
	termbox.Close()
}

// Size return the size of the terminal window
func (*Term) Size() (int, int) {
	return termbox.Size()
}

// Poll gets a terminal event and returns it
func (*Term) Poll() event.Event {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		return parseKeyboardEvent(&ev)
	case termbox.EventMouse:
		return parseMouseEvent
	case termbox.EventResize:
		return &event.Resize{Width: ev.Width, Height: ev.Height}
	case termbox.EventError:
		return &event.Error{Msg: ev.Err.Error()}
	case termbox.EventInterrupt:
		return &event.Quit{}
	}

	return nil
}

// Draw draws a UI element to the terminal window
func (*Term) Draw(item ui.Drawable) {
	buf := ui.NewBuffer(item.GetRect())
	item.Draw(buf)
	for p, cell := range buf.Cells {
		if p.In(buf.Rectangle) {
			fg := termbox.Attribute(cell.Style.FG+1) | termbox.Attribute(cell.Style.Modifier)
			bg := termbox.Attribute(cell.Style.BG + 1)
			termbox.SetCell(p.X, p.Y, cell.Rune, fg, bg)
		}
	}
}

// SetCursor sets the cursor location via termbox
func (*Term) SetCursor(x, y int) {
	termbox.SetCursor(x, y)
}

// Refresh flushes the termbox UI
func (*Term) Refresh() error {
	return termbox.Flush()
}

// Quit sends a quit signal to the window to close the app
func (*Term) Quit() {
	go termbox.Interrupt()
}

func parseKeyboardEvent(ev *termbox.Event) *event.Keyboard {
	var e event.Keyboard
	e.Key = "%s"
	if ev.Mod == termbox.ModAlt {
		e.Key = "<M-%s>"
	}

	if ev.Ch != 0 {
		e.Key = fmt.Sprintf(e.Key, string(ev.Ch))
		return &e
	}

	key, ok := termboxKeys[ev.Key]
	if !ok {
		e.Key = ""
		return &e
	}

	e.Key = fmt.Sprintf(e.Key, key)
	return &e
}

func parseMouseEvent(ev *termbox.Event) *event.Mouse {
	var e event.Mouse
	e.X = ev.MouseX
	e.Y = ev.MouseY
	action, ok := termboxMouseActions[ev.Key]
	if !ok {
		e.Action = event.MouseUnknown
		return &e
	}

	e.Action = action
	return &e
}
