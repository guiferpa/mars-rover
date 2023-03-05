package rover

import "fmt"

type RoverInstructionError struct {
	Letter string
}

func (err *RoverInstructionError) Error() string {
	return fmt.Sprintf("there's a problem with the letter %s", err.Letter)
}

var ErrInvalidInstruction = &RoverInstructionError{""}

type UserCase interface {
	Walk(inst string) error
	String() string
}
