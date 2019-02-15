package govim

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
)

func New(mode gocui.OutputMode, File string) *Gui {
	g, err := gocui.NewGui(mode)
	if err != nil {
		g.Close()
	}

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen
	g.InputEsc = true

	gui := &Gui{Gui: g, File: File}
	gui.init()
	return gui
}

func (g *Gui) init() {
	maxX, maxY := g.Size()
	g.Panels = append(g.Panels, &Panel{Name: "Flist", Body: "File list(work in progress)", X0: -1, Y0: -1, X1: maxX/6 - 1, Y1: maxY})
	g.Panels = append(g.Panels, &Panel{Name: "Status", Body: "StatusBar(work in progress)", X0: maxX / 6, Y0: maxY - 4, X1: maxX, Y1: maxY - 2})
	g.Panels = append(g.Panels, &Panel{Name: "Commands", Body: "CommandsBar(work in progress)", X0: maxX / 6, Y0: maxY - 2, X1: maxX, Y1: maxY})
	g.Panels = append(g.Panels, &Panel{Name: "Numbar", Body: "RowNumber(work in progress)", X0: maxX / 6, Y0: -1, X1: maxX/6 + 4, Y1: maxY - 3})
	g.Panels = append(g.Panels, &Panel{Name: "Editor", Body: g.renderContent(), X0: maxX/6 + 4, Y0: -1, X1: maxX, Y1: maxY - 3, File: FileRead(g.File)})

	for _, panel := range g.Panels {
		if err := panel.SetView(g.Gui); err != nil {
			panic(err)
		}
	}
}

func (gui *Gui) getfiles() error {
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
		gui.File = strings.TrimRight(s, "\n")
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
	File, err := os.Open(gui.File)
	if err != nil {
		return ""
	}
	contents, err := ioutil.ReadAll(File)
	if err != nil {
		return ""
	}
	return string(contents)
}

func (p Panel) SetView(g *gocui.Gui) error {
	if v, err := g.SetView(p.Name, p.X0, p.Y0, p.X1, p.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		if p.File != nil {
			v.Editor = &VimEditor{Insert: false, File: p.File}
		} else {
			v.Editor = &VimEditor{Insert: false}
		}
		v.Editable = true
		switch p.Name {
		case "Editor":
			v.Wrap = true
			if _, err := g.SetCurrentView(p.Name); err != nil {
				return err
			}
			fmt.Fprintln(v, p.Body)
		case "Flist":
			fmt.Fprintln(v, getFiles())
		case "Status":
			fmt.Fprintln(v, "")
		case "Commands":
			fmt.Fprintln(v, p.Body)
			defer v.Clear()
		default:
			fmt.Fprintln(v, p.Body)
		}
		v.Highlight = true
	}
	return nil
}

func getFiles() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ""
	}
	var flist string
	for _, file := range files {
		flist += file.Name() + "\n"
	}
	return flist
}

func FileRead(filename string) *os.File {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	return fp
}
