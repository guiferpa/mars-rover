package rover

import "strings"

type Rover struct {
	dir int
}

func (r *Rover) TurnLeft() {
	if r.dir == East {
		r.dir = North
		return
	}

	r.dir += 1
}

func (r *Rover) Walk(inst string) error {
	for _, l := range strings.Split(inst, "") {
		if _, ok := Letters[l]; !ok {
			ErrInvalidInstruction.Letter = l
			return ErrInvalidInstruction
		}
	}

	return nil
}
