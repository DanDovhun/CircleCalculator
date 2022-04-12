package centre

import (
	"CircleCalculator/funcs"
	"CircleCalculator/structs"
)

func CentreAndCirc(centre, edge structs.Coordinates) structs.Results {
	radius := centre.GetDistance(edge)

	return structs.Results{
		Centre:  centre,
		Radius:  radius,
		Formula: funcs.GetFormula(centre, radius),
	}
}
