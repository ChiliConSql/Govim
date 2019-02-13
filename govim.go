package main

import (
	"fmt"
	govim "govim/pkg"
	"log"
	"os"

	"github.com/jroimartin/gocui"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, fmt.Errorf("File not given\n"))
	}
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	maxX, maxY := g.Size()
	flistpanel := &govim.Panel{Name: "Flist", Body: "file list(work in progress)", X0: -1, Y0: -1, X1: maxX/6 - 1, Y1: maxY}
	statusbar := &govim.Panel{Name: "Status", Body: "StatusBar(work in progress)", X0: maxX / 6, Y0: maxY - 4, X1: maxX, Y1: maxY - 2}
	commandsbar := &govim.Panel{Name: "Commands", Body: "CommandsBar(work in progress)", X0: maxX / 6, Y0: maxY - 2, X1: maxX, Y1: maxY}
	numbar := &govim.Panel{Name: "Numbar", Body: "RowNumber(work in progress)", X0: maxX / 6, Y0: -1, X1: maxX/6 + 4, Y1: maxY - 3}
	editpanel := &govim.Panel{Name: "Editor", Body: "EditPanel(work in progress)", X0: maxX/6 + 4, Y0: -1, X1: maxX, Y1: maxY - 3, File: os.Args[1]}
	g.SetManager(flistpanel, statusbar, commandsbar, numbar, editpanel)
	//g.SetManagerFunc(govim.Layout)

	g.InputEsc = true
	g.Cursor = true
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatal(err)
	}
	if err := g.SetKeybinding("Editor", ':', gocui.ModNone, command); err != nil {
		log.Fatal(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
func command(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView("Commands")
	if err != nil {
		return err
	}
	return nil
}
