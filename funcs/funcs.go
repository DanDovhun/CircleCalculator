package funcs

import (
	"CircleCalculator/structs"
	"fmt"
	"math"
)

func Round(num float64, to int) float64 {
	return math.Round(num*math.Pow10(to)) / math.Pow10(to)
}

func GetFormula(centre structs.Coordinates, radius float64) string {
	var xValue, yValue string

	if centre.X > 0 {
		xValue = fmt.Sprintf("(x + %v)^2", centre.X)
	} else if centre.X < 0 {
		xValue = fmt.Sprintf("(x - %v)^2", -1*centre.X)
	} else {
		xValue = "x^2"
	}

	if centre.Y > 0 {
		yValue = fmt.Sprintf("(y + %v)^2", centre.Y)
	} else if centre.Y < 0 {
		yValue = fmt.Sprintf("(y - %v)^2", -1*centre.Y)
	} else {
		yValue = "x^2"
	}

	return fmt.Sprintf("%v + %v = %v", xValue, yValue, Round(math.Pow(radius, 2), 5))
}
