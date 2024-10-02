package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")

	w.SetContent(widget.NewLabel("TODOs will go here"))
	w.ShowAndRun()
}
