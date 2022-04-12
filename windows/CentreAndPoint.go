package windows

import (
	"CircleCalculator/centre"
	"CircleCalculator/funcs"
	"CircleCalculator/structs"
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func Error(app fyne.App, msg string) {
	w := app.NewWindow("Error")

	w.SetContent(widget.NewLabel(msg))
	w.Show()
}

func CentreAndCirc(app fyne.App, wind fyne.Window) {
	w := app.NewWindow("Circle and Coordinates")

	cnt := widget.NewLabel("")
	radius := widget.NewLabel("")
	formula := widget.NewLabel("")

	w.SetMaster()
	wind.Hide()

	centX := widget.NewEntry()
	centX.SetPlaceHolder("X")

	centY := widget.NewEntry()
	centY.SetPlaceHolder("Y")

	circX := widget.NewEntry()
	circX.SetPlaceHolder("X")

	circY := widget.NewEntry()
	circY.SetPlaceHolder("Y")

	content := container.NewVBox(
		widget.NewLabel("Centre"),
		centX,
		centY,

		widget.NewLabel("Point on a circumference"),
		circX,
		circY,

		widget.NewButton("Submit", func() {
			centXNum, err1 := strconv.ParseFloat(centX.Text, 64)
			centYNum, err2 := strconv.ParseFloat(centY.Text, 64)
			circXNum, err3 := strconv.ParseFloat(circX.Text, 64)
			circYNum, err4 := strconv.ParseFloat(circY.Text, 64)

			if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
				cent := structs.Coordinates{
					X: centXNum,
					Y: centYNum,
				}

				edge := structs.Coordinates{
					X: circXNum,
					Y: circYNum,
				}

				results := centre.CentreAndCirc(cent, edge)

				if cent.GetDistance(edge) == 0 {
					Error(app, "Points cannot be on the same place")
					return
				}

				cnt.SetText(fmt.Sprintf("Centre: %v", results.CoordinatesToString()))
				radius.SetText(fmt.Sprintf("Radius: %v", funcs.Round(results.Radius, 5)))
				formula.SetText(fmt.Sprintf("Formula: %v", results.Formula))
			} else {
				Error(app, "Please only enter numbers")

				return
			}
		}),

		widget.NewButton("Back", func() {
			w.Hide()
			wind.Show()
			wind.SetMaster()
		}),

		cnt,
		radius,
		formula,
	)

	w.Resize(fyne.NewSize(400, 300))
	w.SetContent(content)
	w.Show()
}
