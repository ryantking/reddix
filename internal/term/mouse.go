package term

import (
	"github.com/RyanTKing/reddix/pkg/event"
	termbox "github.com/nsf/termbox-go"
)

var termboxMouseActions = map[termbox.Key]string{
	termbox.MouseLeft:      event.MouseLeft,
	termbox.MouseMiddle:    event.MouseMiddle,
	termbox.MouseRight:     event.MouseRight,
	termbox.MouseRelease:   event.MouseRight,
	termbox.MouseWheelUp:   event.MouseWheelUp,
	termbox.MouseWheelDown: event.MouseWheelDown,
}
