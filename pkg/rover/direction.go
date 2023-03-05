package rover

import "fmt"

func (err *RoverDirectionError) Error() string {
	return fmt.Sprintf("unexpected direction %s", err.Direction)
}

var ErrInvalidDirection = &RoverDirectionError{""}

const (
	North = 0
	West  = 1
	South = 2
	East  = 3
)

type RoverDirectionError struct {
	Direction string
}

var Directions map[int]string

func init() {
	Directions = map[int]string{
		North: "N",
		West:  "W",
		South: "S",
		East:  "E",
	}
}
