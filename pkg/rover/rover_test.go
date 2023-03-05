package rover

import (
	"errors"
	"testing"
)

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

		if got, expected := r.dir, s.Expected; got != expected {
			t.Errorf("unexpected result, got: %v, expected: %v", Directions[got], Directions[expected])
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

func TestSetPosition(t *testing.T) {
	suite := []struct {
		X                 int
		Y                 int
		Direction         string
		ExpectedX         int
		ExpectedY         int
		ExpectedDirection int
	}{
		{1, 2, "N", 1, 2, 0},
		{6, 4, "W", 6, 4, 1},
		{3, 0, "S", 3, 0, 2},
		{5, 5, "E", 5, 5, 3},
	}

	for _, s := range suite {
		r := &Rover{x: 0, y: 0, dir: 0}

		if err := r.SetPosition(s.X, s.Y, s.Direction); err != nil {
			t.Error(err)
			return
		}

		if got, expected := r.x, s.ExpectedX; got != expected {
			t.Errorf("unexpected result for X coords, got: %d, expected: %d", got, expected)
			return
		}

		if got, expected := r.y, s.ExpectedY; got != expected {
			t.Errorf("unexpected result for Y coords, got: %d, expected: %d", got, expected)
			return
		}

		if got, expected := r.dir, s.ExpectedDirection; got != expected {
			t.Errorf("unexpected result for direction, got: %d, expected: %d", got, expected)
			return
		}
	}
}

func TestSetPositionWithInvalidDirection(t *testing.T) {
	suite := []struct {
		Direction string
	}{
		{"A"},
		{"1"},
	}

	for _, s := range suite {
		r := &Rover{x: 0, y: 0, dir: 0}

		got := r.SetPosition(0, 0, s.Direction)
		if !errors.Is(got, ErrInvalidDirection) {
			t.Errorf("unexpected error, got: %v, expected: %v", got, ErrInvalidDirection)
			return
		}

		err := got.(*RoverDirectionError)
		if got, expected := err.Direction, s.Direction; got != expected {
			t.Errorf("unexpected direction reported from error, got: %s, expected: %s", got, expected)
			return
		}
	}
}

func TestGetPosition(t *testing.T) {
	suite := []struct {
		InitialX         int
		InitialY         int
		InitialDirection int
		Expected         string
	}{
		{1, 2, 0, "1 2 N"},
		{6, 4, 1, "6 4 W"},
		{3, 0, 2, "3 0 S"},
		{5, 5, 3, "5 5 E"},
	}

	for _, s := range suite {
		r := &Rover{x: s.InitialX, y: s.InitialY, dir: s.InitialDirection}

		if got, expected := r.GetPosition(), s.Expected; got != expected {
			t.Errorf("unexpected result for get position, got: %s, expected: %s", got, expected)
			return
		}
	}
}
