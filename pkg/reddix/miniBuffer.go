package reddix

import (
	"fmt"
	"strings"

	"github.com/RyanTKing/reddix/pkg/ui"
	"github.com/RyanTKing/reddix/pkg/ui/elements"
	"github.com/RyanTKing/reddix/pkg/window"
)

// NewMiniBuffer creates a new mini buffer
func NewMiniBuffer(y, width int) *MiniBuffer {
	miniBuffer := MiniBuffer{
		Y:     y,
		Width: width,
	}

	return &miniBuffer
}

// GetInput returns the current text input
func (mb *MiniBuffer) GetInput() string {
	return strings.Join(mb.Input, "")
}

// SetStatus sets a status message to the minibuffer
func (mb *MiniBuffer) SetStatus(win window.Window, msg string, err bool) {
	mb.InputMode = false
	mb.Message = msg
	mb.Error = err

	mb.Draw(win)
}

// BeginInput mode enters the minibuffer into input mode
func (mb *MiniBuffer) BeginInput(win window.Window, target string) {
	mb.InputMode = true
	mb.Target = target
	mb.Input = []string{}
	mb.HiddenInput = false
	win.SetCursor(len(mb.Target)+2, mb.Y)
	mb.Draw(win)
}

// Append adds a string of input to the minibuffer
func (mb *MiniBuffer) Append(win window.Window, s string) {
	mb.Input = append(mb.Input, s)
	win.SetCursor(len(mb.Target)+len(mb.GetInput())+2, mb.Y)
	mb.Draw(win)
}

// Backspace removes the last string from the minibuffer
func (mb *MiniBuffer) Backspace(win window.Window) {
	if len(mb.Input) > 0 {
		mb.Input = mb.Input[:len(mb.Input)-1]
		win.SetCursor(len(mb.Target)+len(mb.GetInput())+2, mb.Y)
		mb.Draw(win)
	}
}

// Resize handles a resize event, redrawing if necessary
func (mb *MiniBuffer) Resize(win window.Window, y, width int) {
	if y == mb.Y {
		return
	}

	mb.Y = y
	mb.Width = width
	mb.Draw(win)
}

// Draw draws the mini buffer to the given window
func (mb *MiniBuffer) Draw(win window.Window) {
	mb.fill(win)
	mb.drawText(win)
	win.Refresh()
}

func (mb *MiniBuffer) fill(win window.Window) {
	block := ui.NewBlock()
	block.SetRect(0, mb.Y, mb.Width, mb.Y+1)
	win.Draw(block)
}

func (mb *MiniBuffer) drawText(win window.Window) {
	if mb.InputMode {
		s := mb.GetInput()
		if mb.HiddenInput {
			s = strings.Repeat("*", len(mb.Input))
		}
		text := elements.NewText(fmt.Sprintf("%s: %s", mb.Target, s))
		text.SetRect(0, mb.Y, mb.Width, mb.Y+1)
		win.Draw(text)
		return
	}

	if mb.Error {
		text := elements.NewError(mb.Message)
		text.SetRect(0, mb.Y, mb.Width, mb.Y+1)
		win.Draw(text)
		return
	}

	text := elements.NewText(mb.Message)
	text.SetRect(0, mb.Y, mb.Width, mb.Y+1)
	win.Draw(text)
}
