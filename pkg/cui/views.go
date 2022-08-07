package cui

const IndexView = "index"
const Address = "address"
const Body = "body"

type Position struct {
	// the position relative to the dimension of the terminal
	Relative float32

	// absolute modifier of the position, used to prevent truncation to 0 or
	// to guarantee a certain value if a certain View should always have certain
	// dimensions regardless of the terminal dimensions
	Absolute int
}

type View struct {
	X0 Position
	Y0 Position
	X1 Position
	Y1 Position
}

var Views = map[string]View{
	IndexView: {
		X0: Position{0, 0},
		Y0: Position{0, 0},
		X1: Position{0.15, 0},
		Y1: Position{1, 0},
	},
	Address: {
		X0: Position{0.15, 0},
		Y0: Position{0, 0},
		X1: Position{1, 0},
		Y1: Position{0, 1},
	},
	Body: {
		X0: Position{0.15, 0},
		Y0: Position{0, 1},
		X1: Position{1, 0},
		Y1: Position{1, 0},
	},
}

func (v View) getPositions(maxX, maxY int) (x0 int, y0 int, x1 int, y1 int) {
	x0 = int(float32(maxX)*v.X0.Relative) + v.X0.Absolute
	y0 = int(float32(maxY)*v.Y0.Relative) + v.Y0.Absolute
	x1 = int(float32(maxX)*v.X1.Relative) + v.X1.Absolute
	y1 = int(float32(maxY)*v.Y1.Relative) + v.Y1.Absolute

	return
}
