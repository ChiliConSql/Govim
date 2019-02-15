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
	g := govim.New(gocui.Output256, os.Args[1])
	defer g.Close()

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, govim.Quit); err != nil {
		log.Fatal(err)
	}
	if err := g.SetKeybinding("Editor", ':', gocui.ModNone, govim.Command); err != nil {
		log.Fatal(err)
	}
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
