package robot

import "fmt"

type Position float64
type Facing int
const (
        InvalidFacing Facing = iota
	North
        East
        South
        West
)

var facingName = [...]string{`INVALID`, `NORTH`, `EAST`, `SOUTH`, `WEST`}
var facingId = map[string]Facing{
	`NORTH`: 	North,
	`EAST`:		East,
	`SOUTH`:	South,
	`WEST`:		West,
}

func NewFacing(name string) (Facing, error) {
	f, found := facingId[name]
	if !found {
		return InvalidFacing, fmt.Errorf("Invalid facing: %v", name)
	}
	return f, nil
}

func (f Facing) String() string {
        return facingName[f]
}

type Robot struct {
	minX	Position
	maxX	Position
	minY	Position
	maxY	Position
	X	Position
	Y	Position
	F	Facing
}

func NewRobot(minX, maxX, minY, maxY Position) *Robot {
	return &Robot{
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}
}

func (r *Robot) Place(x, y Position, f Facing) {
	r.X = x
	r.Y = y
	r.F = f
}

func (r *Robot) Move() {
}

func (r *Robot) Dump() {
	fmt.Printf("%+v\n", r)
}
