package windows

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func ThreePoints(app fyne.App, home fyne.Window) {
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

	pntCX := widget.NewEntry()
	pntCX.SetPlaceHolder("X")

	pntCY := widget.NewEntry()
	pntCY.SetPlaceHolder("X")

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Point A"),
		pntAX, pntAY,

		widget.NewLabel("Point B"),
		pntBX, pntBY,

		widget.NewLabel("Point C"),
		pntCX, pntCY,

		widget.NewButton("Calculate", func() {}),

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
