package rover

type RoverSDKRepository interface {
	SetPosition(x, y int, dir string) error
	TurnLeft()
	TurnRight()
	MoveOn()
	GetPosition() string
}

type UserCase interface {
	Walk(inst string) error
	String() string
}
