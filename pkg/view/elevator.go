package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	controller2 "github.com/erbilsilik/elevator-navigation-app/pkg/controller"
)

type Elevator struct {
	elevatorController *controller2.ElevatorController
	output  *widget.Label
	buttons map[string]*widget.Button
	window  fyne.Window
}

func (e *Elevator) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	e.buttons[text] = button
	return button
}

func (e *Elevator) Display(event string, currentFloor string) {
	e.output.SetText(event + " at " + currentFloor)
}

func (e *Elevator) strButton(floor string) *widget.Button {
	return e.addButton(floor, func() {
		e.elevatorController.OnPress(floor)
	})
}

func (e *Elevator) LoadUI(app fyne.App) {
	e.output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	e.output.TextStyle.Monospace = true

	e.window = app.NewWindow("Elevator Navigation")
	e.window.SetContent(container.NewGridWithColumns(1,
		e.output,
		container.NewGridWithColumns(1,
			e.strButton("4")),
		container.NewGridWithColumns(1,
			e.strButton("3")),
		container.NewGridWithColumns(1,
			e.strButton("2")),
		container.NewGridWithColumns(1,
			e.strButton("1")),
		container.NewGridWithColumns(1,
			e.strButton("P1")),
		container.NewGridWithColumns(1,
			e.strButton("P2")),
	))

	e.window.Resize(fyne.NewSize(600, 600))
	e.window.CenterOnScreen()
	e.window.Show()
}

func NewElevator(elevatorController *controller2.ElevatorController) *Elevator {
	return &Elevator{
		elevatorController: elevatorController,
		buttons: make(map[string] * widget.Button, 19),
	}
}
