package rover

import (
	"errors"
	"testing"
)

type RoverSDKMock struct {
	SetPositionError   error
	NCalledTurnLeft    int
	NCalledTurnRight   int
	NCalledMoveOn      int
	NCalledGetPosition int
	GetPositionResult  string
}

func (m *RoverSDKMock) SetPosition(x, y int, dir string) error {
	return m.SetPositionError
}

func (m *RoverSDKMock) TurnLeft() {
	m.NCalledTurnLeft += 1
}

func (m *RoverSDKMock) TurnRight() {
	m.NCalledTurnRight += 1
}

func (m *RoverSDKMock) MoveOn() {
	m.NCalledMoveOn += 1
}

func (m *RoverSDKMock) GetPosition() string {
	m.NCalledGetPosition += 1
	return m.GetPositionResult
}

func TestRoverWalkWithInvalidInstruction(t *testing.T) {
	svc := &UseCaseService{sdk: &RoverSDKMock{}}

	suite := []struct {
		InitialX         int
		InitialY         int
		InitialDirection string
		Instruction      string
	}{
		{0, 1, "N", "LMLMLMLM5"},
		{3, 5, "S", "LMLYLMLMM"},
	}

	for _, s := range suite {
		if _, got := svc.Walk(s.InitialX, s.InitialY, s.InitialDirection, s.Instruction); !errors.Is(got, ErrInvalidInstruction) {
			t.Errorf("unexpected got value different of RoverError struct: got: %v", got)
			return
		}
	}
}

func TestRoverWalk(t *testing.T) {
	suite := []struct {
		InitialX                 int
		InitialY                 int
		InitialDirection         string
		Instruction              string
		ExpectedNCalledTurnLeft  int
		ExpectedNCalledTurnRight int
		ExpectedNCalledMoveOn    int
	}{
		{1, 2, "N", "LMLMLMLMM", 4, 0, 5},
		{3, 3, "E", "MMRMMRMRRM", 0, 4, 6},
		{0, 0, "N", "RMMLMMM", 1, 1, 5},
	}

	for _, s := range suite {
		mock := &RoverSDKMock{}
		svc := &UseCaseService{sdk: mock}

		_, err := svc.Walk(s.InitialX, s.InitialY, s.InitialDirection, s.Instruction)
		if err != nil {
			t.Error(err)
			return
		}

		if got := mock.NCalledTurnLeft; got != s.ExpectedNCalledTurnLeft {
			t.Errorf("unexpected N called turn left, got: %d, expected: %d", got, s.ExpectedNCalledTurnLeft)
			return
		}

		if got := mock.NCalledTurnRight; got != s.ExpectedNCalledTurnRight {
			t.Errorf("unexpected N called turn right, got: %d, expected: %d", got, s.ExpectedNCalledTurnRight)
			return
		}

		if got := mock.NCalledMoveOn; got != s.ExpectedNCalledMoveOn {
			t.Errorf("unexpected N called move on, got: %d, expected: %d", got, s.ExpectedNCalledMoveOn)
			return
		}

		if got, expected := mock.NCalledGetPosition, 1; got != expected {
			t.Errorf("unexpected N called get position, got: %d, expected: %d", got, expected)
			return
		}
	}
}
