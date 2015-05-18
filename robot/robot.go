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

func (r *Robot) Place(x, y Position, f Facing) error {
	if err := r.ValidateBounds(x, y); err != nil {
		return err
	}

	r.X = x
	r.Y = y
	r.F = f

	return nil
}

func (r *Robot) Left() {
	if r.F == North {
		r.F = West
	} else {
		r.F--
	}
}

func (r *Robot) Right() {
	if r.F == West {
		r.F = North
	} else {
		r.F++
	}
}

func (r *Robot) Move() error {
	newX, newY := r.X, r.Y
	switch r.F {
	case North:
		newY++
	case East:
		newX++
	case South:
		newY--
	case West:
		newX--
	}

	if err := r.ValidateBounds(newX, newY); err != nil {
		return err
	}

	r.X = newX
	r.Y = newY
	return nil
}

func (r *Robot) ValidateBounds(x,y Position) error {
	if r.minX <= x && x <= r.maxX &&
	   r.minY <= y && y <= r.maxY {
		return nil
	}
	return fmt.Errorf("Position %v,%v out of bounds", x, y)
}

func (r *Robot) Report() {
	fmt.Printf("%v,%v,%v\n", r.X, r.Y, r.F)
}

func (r *Robot) Dump() {
	fmt.Printf("%+v\n", r)
}
