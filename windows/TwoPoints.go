package windows

import (
	"CircleCalculator/opposites"
	"CircleCalculator/structs"
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func TwoPoints(app fyne.App, home fyne.Window) {
	w := app.NewWindow("Circle and Coordinates")

	cnt := widget.NewLabel("")
	radius := widget.NewLabel("")
	formula := widget.NewLabel("")

	w.SetMaster()
	home.Hide()

	pntAX := widget.NewEntry()
	pntAX.SetPlaceHolder("X")

	pntAY := widget.NewEntry()
	pntAY.SetPlaceHolder("Y")

	pntBX := widget.NewEntry()
	pntBX.SetPlaceHolder("X")

	pntBY := widget.NewEntry()
	pntBY.SetPlaceHolder("Y")

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Point A"),
		pntAX,
		pntAY,

		widget.NewLabel("Point B"),
		pntBX,
		pntBY,

		widget.NewButton("Calculate", func() {
			pntAXNum, err1 := strconv.ParseFloat(pntAX.Text, 64)
			pntAYNum, err2 := strconv.ParseFloat(pntAY.Text, 64)
			pntBXNum, err3 := strconv.ParseFloat(pntBX.Text, 64)
			pntBYNum, err4 := strconv.ParseFloat(pntBY.Text, 64)

			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				Error(app, "Please only enter numbers")
				return
			}

			pointA := structs.Coordinates{
				X: pntAXNum,
				Y: pntAYNum,
			}

			pointB := structs.Coordinates{
				X: pntBXNum,
				Y: pntBYNum,
			}

			if pointA.GetDistance(pointB) == 0 {
				Error(app, "Points cannot be on the same place")
				return
			}

			res := opposites.Opposites(pointA, pointB)

			cnt.SetText(fmt.Sprintf("Results: %v", res.CoordinatesToString()))
			radius.SetText(fmt.Sprintf("Radius: %v", res.Radius))
			formula.SetText(fmt.Sprintf("Formula: %v", res.Formula))
		}),

		widget.NewButton("Back", func() {
			w.Hide()
			home.Show()
			home.SetMaster()
		}),

		cnt,
		radius,
		formula,
	))

	w.Resize(fyne.NewSize(400, 300))
	w.Show()
}
