package window

import (
	"github.com/RyanTKing/reddix/internal/ui/elements"
	termbox "github.com/nsf/termbox-go"
)

const (
	topMenuLeft     = "q:quit r:subreddit u:user h:help"
	topMenuCenter   = "reddix"
	topMenuRight1   = "l:login"
	topMenuRight2   = "l:logout"
	bottomMenuRight = "Not Logged In"
)

func makeTopMenu(width int) *elements.Menu {
	menu := elements.NewMenu(topMenuLeft, topMenuCenter, topMenuRight1)
	menu.SetRect(0, 0, width, 1)
	return menu
}

func makeBottomMenu(width, height int) *elements.Menu {
	menu := elements.NewMenu(frontpage, "", bottomMenuRight)
	menu.SetRect(0, height-2, width, height-1)
	return menu
}

func (win *Window) drawTopMenu() {
	win.draw(win.TopMenu)
	termbox.Flush()
}

func (win *Window) drawBottomMenu() {
	win.draw(win.BottomMenu)
	termbox.Flush()
}
