package server

import (
	"errors"
	"fmt"
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/zkc/pkg/zk"
)

type Server struct {
	kasten *zk.Kasten
	g      *gocui.Gui
}

func NewServer() (*Server, error) {
	g, err := gocui.NewGui(gocui.Output256, true)
	if err != nil {
		return nil, err
	}

	g.Cursor = true

	m1 := gocui.ManagerFunc(func(g *gocui.Gui) error {
		maxX, maxY := g.Size()
		if v, err := g.SetView("hello", 0, 0, maxX/8-7, maxY, 0); err != nil {
			if !errors.Is(err, gocui.ErrUnknownView) {
				return err
			}

			v.Highlight = true
			v.SelBgColor = gocui.ColorGreen
			v.SelFgColor = gocui.ColorBlack

			if _, err := g.SetCurrentView("hello"); err != nil {
				return err
			}

			fmt.Fprintln(v, "Hello world!")
		}

		return nil
	})

	g.SetManager(m1)

	if v, err := g.View("hello"); err == nil {
		v.Editor = gocui.DefaultEditor
	}

	// g.Update(func(g *gocui.Gui) error {
	// 	v, err := g.View("hello")
	// 	if err != nil {
	// 		log.Panic("broke in update; no idea why")
	// 	}
	// 	v.Clear()
	// 	fmt.Fprintln(v, time.Now())
	// 	return nil
	// })

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	server := Server{
		g: g,
	}

	return &server, nil
}

func (s *Server) Start() error {
	defer s.g.Close()

	return s.g.MainLoop()
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
