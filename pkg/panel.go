package govim

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jroimartin/gocui"
)

type Panel struct {
	Name           string
	Body           string
	X0, Y0, X1, Y1 int
	File           string
}

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
			//info,err := os.Lstat(p.File)
			file, err := os.Open(p.File)
			if err != nil {
				return err
			}
			if contents, err := ioutil.ReadAll(file); err != nil {
				return err
			} else {
				fmt.Fprintln(v, string(contents))
			}
		case "Flist":
			dir, err := os.Getwd()
			if err != nil {
				return err
			}
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				return err
			}
			var flist string
			for _, file := range files {
				flist += file.Name() + "\n"
			}
			fmt.Fprintln(v, flist)
		case "Status":
		default:
			fmt.Fprintln(v, p.Body)
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
	}
	return nil
}
