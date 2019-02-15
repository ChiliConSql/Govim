package govim

import (
	"log"
	"os"
	"strings"

	"github.com/jroimartin/gocui"
)

type VimEditor struct {
	Insert bool
	File   *os.File
}

func (ve *VimEditor) Edit(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	if ve.Insert {
		ve.InsertMode(v, key, ch, mod)
	} else {
		ve.NormalMode(v, key, ch, mod)
	}
}

func (ve *VimEditor) InsertMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	if v.Name() == "Commands" && key == gocui.KeyEnter {
		ve.parseCommands(v, key, ch, mod)
		v.Clear()
		ve.Insert = false
	}
	switch {
	case key == gocui.KeyEsc:
		ve.Insert = false
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyDelete:
		v.EditDelete(false)
	case key == gocui.KeyInsert:
		v.Overwrite = !v.Overwrite
	case key == gocui.KeyEnter:
		v.EditNewLine()
	case key == gocui.KeyArrowDown:
		v.MoveCursor(0, 1, false)
	case key == gocui.KeyArrowUp:
		v.MoveCursor(0, -1, false)
	case key == gocui.KeyArrowLeft:
		v.MoveCursor(-1, 0, false)
	case key == gocui.KeyArrowRight:
		v.MoveCursor(1, 0, false)
	}
	// TODO: handle other keybindings...
}

func (ve *VimEditor) NormalMode(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch ch {
	case 'x':
		v.EditDelete(false)
	case 'i':
		ve.Insert = true
	case 'o':
		v.EditNewLine()
		ve.Insert = true
	case 'O':
		v.MoveCursor(0, -1, false)
		v.EditNewLine()
		ve.Insert = true
	case 'j':
		v.MoveCursor(0, 1, false)
	case 'k':
		v.MoveCursor(0, -1, false)
	case 'h':
		v.MoveCursor(-1, 0, false)
	case 'l':
		v.MoveCursor(1, 0, false)
	case ':':
		ve.Insert = true
	}
	// TODO: handle other keybindings...
}

func (ve *VimEditor) parseCommands(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	strings.Map(func(r rune) rune {
		switch r {
		case 'w':

		case 'q':
			log.Fatal(gocui.ErrQuit)
		}
		return r
	}, v.ViewBuffer())
}
