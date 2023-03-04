package rover

import "strings"

type Rover struct{}

func (r *Rover) Walk(inst string) error {
	for _, l := range strings.Split(inst, "") {
		if _, ok := Letters[l]; !ok {
			ErrInvalidInstruction.Letter = l
			return ErrInvalidInstruction
		}
	}

	return nil
}
