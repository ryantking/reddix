package main

import (
	"log"
	"unicode"

	"github.com/jroimartin/gocui"
	// "github.com/jzelinskie/geddit"
)

// setKey sets a key (upercase or lowercase) with a specified function that
// applies to any view
func setKey(g *gocui.Gui, key rune, handler func(*gocui.Gui, *gocui.View) error) {
	lowcase := unicode.ToLower(key)
	upcase := unicode.ToUpper(key)
	if err := g.SetKeybinding("", lowcase, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", upcase, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
}

func main() {
	// Initialize GUI
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// Layout & Keybindings
	g.SetManagerFunc(frontpage)

	setKey(g, 'q', quit)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	// Exit
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
