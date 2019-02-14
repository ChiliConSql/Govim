package govim

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jroimartin/gocui"
)

func (p Panel) Layout(g *gocui.Gui) error {
	if v, err := g.SetView(p.Name, p.X0, p.Y0, p.X1, p.Y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Editor = &VimEditor{Insert: false}
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
