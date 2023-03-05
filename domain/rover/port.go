package rover

type RoverSDKRepository interface {
	SetPosition(x, y int, dir string) error
	TurnLeft()
	TurnRight()
	MoveOn()
	GetPosition() string
}

type UserCase interface {
	Walk(x, y int, dir, inst string) (string, error)
}
