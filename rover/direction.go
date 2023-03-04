package rover

const (
	North = 0
	West  = 1
	South = 2
	East  = 3
)

var Directions map[int]string

func init() {
	Directions = map[int]string{
		North: "N",
		West:  "W",
		South: "S",
		East:  "E",
	}
}
