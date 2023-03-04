package rover

import (
	"errors"
	"testing"
)

func TestInvalidInstructionToWalk(t *testing.T) {
	r := &Rover{}

	suite := []struct {
		Param string
	}{
		{"LMLMLMLM5"},
		{"LMLYLMLMM"},
	}

	for _, s := range suite {
		if got := r.Walk(s.Param); !errors.Is(got, ErrInvalidInstruction) {
			t.Errorf("unexpected got value different of RoverError struct: got: %v", got)
			return
		}
	}
}

func TestRoverTurnLeft(t *testing.T) {
	suite := []struct {
		InitialDirection int
		Turn             int
		Expected         int
	}{
		{North, 4, North},
		{North, 5, West},
		{North, 1, West},
		{North, 2, South},
	}

	for _, s := range suite {
		r := &Rover{s.InitialDirection}

		for s.Turn > 0 {
			s.Turn -= 1
			r.TurnLeft()
		}

		if got := r.dir; got != s.Expected {
			t.Errorf("unexpected result, got: %v, expected: %v", Directions[got], Directions[s.Expected])
			return
		}
	}
}
