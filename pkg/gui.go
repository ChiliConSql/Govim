package govim

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
)

func New(mode gocui.OutputMode, file string) *Gui {
	g, err := gocui.NewGui(mode)
	if err != nil {
		g.Close()
	}

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen
	g.InputEsc = true

	gui := &Gui{Gui: g, file: file}
	gui.init()
	return gui
}

func (g *Gui) init() {
	maxX, maxY := g.Size()
	flistpanel := &Panel{Name: "Flist", Body: "file list(work in progress)", X0: -1, Y0: -1, X1: maxX/6 - 1, Y1: maxY}
	statusbar := &Panel{Name: "Status", Body: "StatusBar(work in progress)", X0: maxX / 6, Y0: maxY - 4, X1: maxX, Y1: maxY - 2}
	commandsbar := &Panel{Name: "Commands", Body: "CommandsBar(work in progress)", X0: maxX / 6, Y0: maxY - 2, X1: maxX, Y1: maxY}
	numbar := &Panel{Name: "Numbar", Body: "RowNumber(work in progress)", X0: maxX / 6, Y0: -1, X1: maxX/6 + 4, Y1: maxY - 3}
	editpanel := &Panel{Name: "Editor", Body: g.renderContent(), X0: maxX/6 + 4, Y0: -1, X1: maxX, Y1: maxY - 3}
	g.SetManager(flistpanel, statusbar, commandsbar, numbar, editpanel)
}

func (gui *Gui) getFile() error {
	gui.Update(func(g *gocui.Gui) error {
		v, err := g.View("Flist")
		if err != nil {
			return err
		}
		_, y := v.Cursor()
		s, err := v.Line(y)
		if err != nil {
			return err
		}
		gui.file = strings.TrimRight(s, "\n")
		vi, err := g.SetCurrentView("Editor")
		if err != nil {
			return err
		}
		vi.Clear()
		fmt.Fprintln(vi, gui.renderContent())
		return nil
	})
	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func Command(g *gocui.Gui, v *gocui.View) error {
	_, err := g.SetCurrentView("Commands")
	if err != nil {
		return err
	}
	return nil
}

func (gui *Gui) renderContent() string {
	file, err := os.Open(gui.file)
	if err != nil {
		return ""
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(contents)
}
