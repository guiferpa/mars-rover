package rover

import (
	"fmt"
	"strings"
)

type RoverInstructionError struct {
	Letter string
}

func (err *RoverInstructionError) Error() string {
	return fmt.Sprintf("there's a problem with the letter %s", err.Letter)
}

var ErrInvalidInstruction = &RoverInstructionError{""}

const (
	TurnLeft  = "L"
	TurnRight = "R"
	MoveOn    = "M"
)

type UseCaseService struct {
	sdk RoverSDKRepository
}

func (s *UseCaseService) Walk(x, y int, dir string, inst string) (string, error) {
	if err := s.sdk.SetPosition(x, y, dir); err != nil {
		return "", err
	}

	for _, letter := range strings.Split(inst, "") {
		switch letter {
		case TurnLeft:
			s.sdk.TurnLeft()

		case TurnRight:
			s.sdk.TurnRight()

		case MoveOn:
			s.sdk.MoveOn()

		default:
			ErrInvalidInstruction.Letter = letter
			return "", ErrInvalidInstruction
		}
	}

	return s.sdk.GetPosition(), nil
}

func NewUseCaseService(sdk RoverSDKRepository) *UseCaseService {
	return &UseCaseService{sdk}
}
