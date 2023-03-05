package rover

import (
	"strings"
)

const (
	TurnLeft  = "L"
	TurnRight = "R"
	MoveOn    = "M"
)

type Service struct {
	sdk RoverSDKRepository
}

func (s *Service) Walk(x, y int, dir string, inst string) (string, error) {
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

func NewService(sdk RoverSDKRepository) *Service {
	return &Service{sdk}
}
