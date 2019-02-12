package govim

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

type Panel struct {
	Name           string
	Body           string
	X0, Y0, X1, Y1 int
}

func (p Panel) Layout(g *gocui.Gui) error {
	if v, err := g.SetView(p.Name, p.X0, p.Y0, p.X1, p.Y1); err != nil {
		v.Highlight = true
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, p.Body)
	}
	return nil
}

/*
func NewGui() (*Gui, error) {
	file, err := os.Open("example.txt")
	if err != nil {
		return nil, err
	}
	f, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &Gui{mode: &Mode{mode: "Normal", file: f, selectedLine: 0}}, nil
}
*/
