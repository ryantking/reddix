package main

import (
	"log"

	"github.com/jroimartin/gocui"
	// "github.com/jzelinskie/geddit"
)

func main() {
	// Initialize GUI
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	// Layout & Keybindings
	g.SetManagerFunc(frontpage)

	err = g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit)
	if err != nil {
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
