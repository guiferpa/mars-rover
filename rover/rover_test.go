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
		r := &Rover{s.InitialDirection, 0, 0}

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

func TestRoverTurnRight(t *testing.T) {
	suite := []struct {
		InitialDirection int
		Turn             int
		Expected         int
	}{
		{North, 4, North},
		{North, 5, East},
		{South, 5, West},
	}

	for _, s := range suite {
		r := &Rover{s.InitialDirection, 0, 0}

		for s.Turn > 0 {
			s.Turn -= 1
			r.TurnRight()
		}

		if got := r.dir; got != s.Expected {
			t.Errorf("unexpected result, got: %v, expected: %v", Directions[got], Directions[s.Expected])
			return
		}
	}
}

func TestRoverMoveOn(t *testing.T) {
	suite := []struct {
		InitialDirection int
		Turn             int
		ExpectedX        int
		ExpectedY        int
	}{
		{North, 3, 0, 3},
		{West, 2, -2, 0},
		{South, 1, 0, -1},
		{East, 8, 8, 0},
	}

	for _, s := range suite {
		r := &Rover{s.InitialDirection, 0, 0}

		for s.Turn > 0 {
			s.Turn -= 1
			r.MoveOn()
		}

		if got := r.x; got != s.ExpectedX {
			t.Errorf("unexpected result for coord X, got: %v, expected: %v", got, s.ExpectedX)
			return
		}

		if got := r.y; got != s.ExpectedY {
			t.Errorf("unexpected result for coord Y, got: %v, expected: %v", got, s.ExpectedY)
			return
		}
	}
}
