package rover

import "fmt"

const (
	TurnLeft  = "L"
	TurnRight = "R"
	MoveOn    = "M"
)

var Letters map[string]bool

type RoverInstructionError struct {
	Letter string
}

func (err *RoverInstructionError) Error() string {
	return fmt.Sprintf("there's a problem with the letter %s", err.Letter)
}

var (
	ErrInvalidInstruction = &RoverInstructionError{""}
)

func init() {
	Letters = map[string]bool{
		TurnLeft:  true,
		TurnRight: true,
		MoveOn:    true,
	}
}
