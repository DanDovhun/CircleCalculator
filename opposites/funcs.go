package opposites

import (
	"CircleCalculator/funcs"
	"CircleCalculator/structs"
)

func Opposites(pointA, pointB structs.Coordinates) structs.Results {
	centre := structs.Coordinates{
		X: (pointA.X + pointB.X) / 2,
		Y: (pointA.Y + pointB.Y) / 2,
	}

	radius := pointA.GetDistance(centre)

	return structs.Results{
		Centre:  centre,
		Radius:  radius,
		Formula: funcs.GetFormula(centre, radius),
	}
}
