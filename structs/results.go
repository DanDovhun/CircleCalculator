package structs

import "fmt"

type Results struct {
	Centre  Coordinates
	Radius  float64
	Formula string
}

func (res Results) CoordinatesToString() string {
	return fmt.Sprintf("(%v, %v)", res.Centre.X, res.Centre.Y)
}
