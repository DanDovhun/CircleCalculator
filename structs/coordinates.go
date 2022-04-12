package structs

import "math"

type Coordinates struct {
	X float64
	Y float64
}

func (a Coordinates) GetDistance(b Coordinates) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func (a Coordinates) SamePosition(b Coordinates) bool {
	if a.X == b.X || a.Y == b.Y {
		return true
	}

	return false
}
