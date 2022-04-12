package main

import (
	"CircleCalculator/windows"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Circle Calculator")

	label := widget.NewLabel("Calculate from: ")

	content := container.NewVBox(
		label,
		widget.NewButton("Centre and a coordinate on circumference", func() {
			windows.CentreAndCirc(a, w)
		}),

		widget.NewButton("Two opposing coordinates", func() {
			windows.TwoPoints(a, w)
		}),

		widget.NewButton("Three random points", func() {
			windows.ThreePoints(a, w)
		}),
	)

	w.SetContent(content)
	w.ShowAndRun()
}
