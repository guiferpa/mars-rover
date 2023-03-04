package rover

import (
	"fmt"
	"strings"
)

type Rover struct {
	dir int
	x   int
	y   int
}

func (r *Rover) TurnLeft() {
	if r.dir == East {
		r.dir = North
		return
	}

	r.dir += 1
}

func (r *Rover) TurnRight() {
	if r.dir == North {
		r.dir = East
		return
	}

	r.dir -= 1
}

func (r *Rover) MoveOn() {
	if r.dir == North {
		r.y += 1
		return
	}

	if r.dir == West {
		r.x -= 1
		return
	}

	if r.dir == South {
		r.y -= 1
		return
	}

	if r.dir == East {
		r.x += 1
		return
	}
}

func (r *Rover) Walk(inst string) error {
	for _, l := range strings.Split(inst, "") {
		if _, ok := Letters[l]; !ok {
			ErrInvalidInstruction.Letter = l
			return ErrInvalidInstruction
		}

		if l == TurnLeft {
			r.TurnLeft()
			continue
		}

		if l == TurnRight {
			r.TurnRight()
			continue
		}

		if l == MoveOn {
			r.MoveOn()
			continue
		}
	}

	return nil
}

func (r *Rover) String() string {
	return fmt.Sprintf("%d %d %s", r.x, r.y, Directions[r.dir])
}
