package cui

import (
	"errors"
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/zkc/pkg/zk"
)

var outputs []gocui.OutputMode = []gocui.OutputMode{
	gocui.Output256,
	gocui.Output216,
	gocui.OutputNormal,
	gocui.OutputGrayscale,
}

type Server struct {
	kasten *zk.Kasten
	g      *gocui.Gui
}

func NewServer() (*Server, error) {
	var g *gocui.Gui
	var err error
	for _, output := range outputs {
		g, err = gocui.NewGui(output, true)
		if err == nil {
			// TODO: display output selection
			break
		}
	}
	if err != nil {
		// TODO: display failure message
		return nil, err
	}

	g.Cursor = true

	g.SetManager(gocui.ManagerFunc(Manager))

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return nil, err
	}

	if v, err := g.View(IndexView); err == nil {
		v.Editor = gocui.DefaultEditor
	}

	return &Server{
		kasten: nil,
		g:      g,
	}, nil
}

func Manager(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	for name, view := range Views {
		x0, y0, x1, y1 := view.getPositions(maxX, maxY)
		// we don't need to overlap any views, so hardcode overlaps to 0
		if v, err := g.SetView(name, x0, y0, x1, y1, 0); err != nil {
			if !errors.Is(err, gocui.ErrUnknownView) {
				return err
			}
			if _, err := g.SetCurrentView(name); err != nil {
				return err
			}
			fmt.Fprintln(v, name)
		}

	}
	return nil
}

func (s *Server) Start() error {
	defer s.g.Close()

	return s.g.MainLoop()
}
