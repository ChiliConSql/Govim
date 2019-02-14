package govim

import "github.com/jroimartin/gocui"

type Gui struct {
	*gocui.Gui
	file string
}

type Panel struct {
	Name           string
	Body           string
	X0, Y0, X1, Y1 int
}
