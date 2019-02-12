package main

import (
	govim "govim/pkg"
	"log"

	"github.com/jroimartin/gocui"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	maxX, maxY := g.Size()
	flistpanel := &govim.Panel{Name: "Flist", Body: "file list(work in progress)", X0: -1, Y0: -1, X1: maxX/6 - 1, Y1: maxY}
	statusbar := &govim.Panel{Name: "Status", Body: "StatusBar(work in progress)", X0: maxX / 6, Y0: maxY - 3, X1: maxX, Y1: maxY}
	numbar := &govim.Panel{Name: "Numbar", Body: "RowNumber(work in progress)", X0: maxX / 6, Y0: -1, X1: maxX/6 + 4, Y1: maxY - 3}
	editpanel := &govim.Panel{Name: "Editor", Body: "EditPanel(work in progress)", X0: maxX/6 + 4, Y0: -1, X1: maxX, Y1: maxY - 3}
	g.SetManager(flistpanel, statusbar, numbar, editpanel)
	//g.SetManagerFunc(govim.Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
