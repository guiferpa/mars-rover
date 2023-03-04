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
