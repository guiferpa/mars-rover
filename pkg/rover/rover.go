package rover

import (
	"fmt"
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

func (r *Rover) SetPosition(x int, y int, dir string) error {
	r.x = x
	r.y = y

	switch dir {
	case Directions[North]:
		r.dir = North

	case Directions[West]:
		r.dir = West

	case Directions[South]:
		r.dir = South

	case Directions[East]:
		r.dir = East

	default:
		ErrInvalidDirection.Direction = dir
		return ErrInvalidDirection
	}

	return nil
}

func (r *Rover) GetPosition() string {
	return fmt.Sprintf("%d %d %s", r.x, r.y, Directions[r.dir])
}
