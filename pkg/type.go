package govim

import (
	"os"

	"github.com/jroimartin/gocui"
)

type Gui struct {
	*gocui.Gui
	File   string
	Panels []*Panel
}

type Panel struct {
	Name           string
	Body           string
	X0, Y0, X1, Y1 int
	File           *os.File
}
