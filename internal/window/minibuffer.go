package window

import (
	"fmt"
	"strings"

	"github.com/RyanTKing/reddix/internal/ui/elements"
	termbox "github.com/nsf/termbox-go"
)

func makeMiniBuffer(width, height int) *MiniBuffer {
	miniBuffer := MiniBuffer{
		Text: elements.NewText(),
	}
	miniBuffer.Text.SetRect(0, height-1, width, height)

	return &miniBuffer
}

func (win *Window) setStatus(msg string, err bool) {
	win.miniBuffer.InputMode = false
	win.miniBuffer.Message = msg
	win.miniBuffer.Error = err

	win.drawMiniBuffer()
}

func (win *Window) enterInputMode(target string) {
	win.mode = TextEntry
	win.miniBuffer.InputMode = true
	win.miniBuffer.Target = target
	win.miniBuffer.Input = ""
	win.miniBuffer.HiddenInput = false
	win.miniBuffer.Error = false
	termbox.SetCursor(len(win.miniBuffer.Target)+2, win.Height-1)

	win.drawMiniBuffer()
}

func (win *Window) exitInputMode() {
	win.mode = Browse
	termbox.SetCursor(0, 0)
	termbox.Flush()
}

func (win *Window) miniBufferAppendChar(r rune) {
	win.miniBuffer.Input += string(r)
	termbox.SetCursor(len(win.miniBuffer.Target)+len(win.miniBuffer.Input)+2, win.Height-1)

	win.drawMiniBuffer()
}

func (win *Window) miniBufferBackspace() {
	if len(win.miniBuffer.Input) > 0 {
		win.miniBuffer.Input = win.miniBuffer.Input[:len(win.miniBuffer.Input)-1]
		termbox.SetCursor(len(win.miniBuffer.Target)+len(win.miniBuffer.Input)+2, win.Height-1)

		win.drawMiniBuffer()
	}
}

func (win *Window) drawMiniBuffer() {
	win.miniBuffer.Text.Error = win.miniBuffer.Error
	if win.miniBuffer.InputMode {
		input := win.miniBuffer.Input
		if win.miniBuffer.HiddenInput {
			input = strings.Repeat("*", len(input))
		}
		win.miniBuffer.Text.Text = fmt.Sprintf("%s: %s", win.miniBuffer.Target, input)
	} else {
		win.miniBuffer.Text.Text = win.miniBuffer.Message
	}

	win.draw(win.miniBuffer.Text)

	termbox.Flush()
}
