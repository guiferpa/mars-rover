package rover

import roverpkg "github/guiferpa/mars-rover/pkg/rover"

type RoverSDK struct {
	rover *roverpkg.Rover
}

func (rs *RoverSDK) SetPosition(x, y int, dir string) error {
	return rs.rover.SetPosition(x, y, dir)
}

func (rs *RoverSDK) TurnLeft() {
	rs.rover.TurnLeft()
}

func (rs *RoverSDK) TurnRight() {
	rs.rover.TurnRight()
}

func (rs *RoverSDK) MoveOn() {
	rs.rover.MoveOn()
}

func (rs *RoverSDK) GetPosition() string {
	return rs.rover.GetPosition()
}

func NewRoverSDK() *RoverSDK {
	return &RoverSDK{
		rover: &roverpkg.Rover{},
	}
}
