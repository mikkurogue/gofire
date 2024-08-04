package ui

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ChatWindow struct {
	Title  string
	window fyne.Window
}

func OpenChatWindow(app fyne.App, title string) *ChatWindow {

	window := app.NewWindow(title)
	window.Resize(fyne.NewSize(float32(500), float32(500)))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Type your message here...")

	exampleReceived := canvas.NewText("Want to play wow?", color.White)

	exampleReceived.TextStyle.Monospace = true

	exampleSent := canvas.NewText("yh lets do some m+", color.White)

	exampleSent.TextStyle.Monospace = true

	receivedMessage := container.New(layout.NewHBoxLayout(), exampleReceived)

	sentMessage := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), exampleSent)

	inputAndSend := container.NewVBox(input, widget.NewButton("Send", func() {
		log.Println("content was:", input.Text)
	}))

	window.SetContent(container.New(layout.NewVBoxLayout(), receivedMessage, sentMessage, inputAndSend))

	window.Show()

	return &ChatWindow{
		Title:  title,
		window: window,
	}
}

func (c *ChatWindow) UpdateChatHistory() {
	// TODO: Create this function
}

func (c *ChatWindow) SendMessage() {
	// TODO: Create this function
}

func (c *ChatWindow) CloseChatWindow() {
	c.window.Close()
}
