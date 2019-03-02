package view

import (
	"github.com/RyanTKing/reddix/pkg/ui"
	"github.com/RyanTKing/reddix/pkg/window"
)

func clear(view View, win window.Window) {
	block := ui.NewBlock()
	block.SetRect(view.Rect())
	win.Draw(block)
}
