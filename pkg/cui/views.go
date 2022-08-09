package cui

import "github.com/awesome-gocui/gocui"

const (
	Index = "Index"
	Title = "Title"
	Body  = "Body"
	Meta  = "Meta"
)

type Position struct {
	// the position relative to the dimension of the terminal
	Relative float32

	// absolute modifier of the position, used to prevent truncation to 0 or
	// to guarantee a certain value if a certain View should always have certain
	// dimensions regardless of the terminal dimensions
	Absolute int
}

type View struct {
	Title    string
	X0       Position
	Y0       Position
	X1       Position
	Y1       Position
	Overlaps byte
}

var Views = map[string]View{
	Index: {
		Title:    Index,
		X0:       Position{0, 0},
		Y0:       Position{0, 0},
		X1:       Position{0.15, 0},
		Y1:       Position{1, -1},
		Overlaps: gocui.RIGHT,
	},
	Title: {
		Title:    Title,
		X0:       Position{0.15, 0},
		Y0:       Position{0, 0},
		X1:       Position{1, -1},
		Y1:       Position{0, 2},
		Overlaps: gocui.BOTTOM,
	},
	Body: {
		Title:    Body,
		X0:       Position{0.15, 0},
		Y0:       Position{0, 2},
		X1:       Position{1, -1},
		Y1:       Position{1, -3},
		Overlaps: gocui.TOP | gocui.BOTTOM,
	},
	Meta: {
		Title:    Meta,
		X0:       Position{0.15, 0},
		Y0:       Position{1, -3},
		X1:       Position{1, -1},
		Y1:       Position{1, -1},
		Overlaps: gocui.TOP,
	},
}

func (v View) getPositions(maxX, maxY int) (x0 int, y0 int, x1 int, y1 int) {
	x0 = int(float32(maxX)*v.X0.Relative) + v.X0.Absolute
	y0 = int(float32(maxY)*v.Y0.Relative) + v.Y0.Absolute
	x1 = int(float32(maxX)*v.X1.Relative) + v.X1.Absolute
	y1 = int(float32(maxY)*v.Y1.Relative) + v.Y1.Absolute

	return
}
