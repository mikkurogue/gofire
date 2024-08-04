package ui

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Window struct {
	Title  string
	Size   [2]int32
	window fyne.Window
	label  *widget.Label
	app    fyne.App
}

// CreateWindow initializes and returns a new Window instance
func CreateWindow(title string, width, height int32) *Window {
	a := app.New()
	w := a.NewWindow(title)
	w.Resize(fyne.NewSize(float32(width), float32(height)))

	window := &Window{
		Title:  title,
		Size:   [2]int32{width, height},
		window: w,
		app:    a,
		label:  widget.NewLabel("Initializing..."),
	}

	w.SetContent(container.NewVBox(
		window.label,
	))

	return window
}

func (w *Window) Show() {
	go w.updateLoop()
	w.window.ShowAndRun()
}

func (w *Window) UpdateLabel(text string) {
	w.label.SetText(text)
}

func (w *Window) updateLoop() {
	// TODO: Figuure out what to update in this, we need to update everything probably whenever a state changes
}

// KillWindow gracefully shuts down the application
func (w *Window) KillWindow() {
	w.app.Quit()
	os.Exit(1)
}
