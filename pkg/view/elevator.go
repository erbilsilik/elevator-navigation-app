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

func (e *Elevator) internalRequestButton(floor string) *widget.Button {
	return e.addButton(floor, func() {
		e.elevatorController.OnPress(floor, "")
	})
}

func (e *Elevator) externalRequestButton(destinationFloor string, sourceFloor string) *widget.Button {
	return e.addButton(destinationFloor, func() {
		if destinationFloor == "Up" {
			destinationFloor = "6"
		} else if destinationFloor == "Down" {
			destinationFloor = "1"
		}
		e.elevatorController.OnPress(destinationFloor, sourceFloor)
	})
}

func (e *Elevator) LoadUI(app fyne.App) {
	e.output = &widget.Label{Alignment: fyne.TextAlignTrailing}
	e.output.TextStyle.Monospace = true

	e.window = app.NewWindow("Elevator Navigation")

	e.window.SetContent(container.NewGridWithColumns(1,
		widget.NewLabel("Inside buttons"),
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(7,
				e.internalRequestButton("1"),
				e.internalRequestButton("2"),
				e.internalRequestButton("3"),
				e.internalRequestButton("4"),
				e.internalRequestButton("5"),
				e.internalRequestButton("6"),
			),
		),
		widget.NewSeparator(),
		widget.NewLabel("Outside buttons"),
		container.NewGridWithColumns(6,
			widget.NewLabel("Floor: 6"),
			e.externalRequestButton("1", "6"),
			e.externalRequestButton("2", "6"),
			e.externalRequestButton("3", "6"),
			e.externalRequestButton("4", "6"),
			e.externalRequestButton("5", "6"),
		),
		container.NewGridWithColumns(3,
			widget.NewLabel("Floor: 5"),
			e.externalRequestButton("Up", "5"),
			e.externalRequestButton("Down", "5"),
		),
		container.NewGridWithColumns(6,
			widget.NewLabel("Floor: 4"),
			e.externalRequestButton("1", "4"),
			e.externalRequestButton("2", "4"),
			e.externalRequestButton("3", "4"),
			e.externalRequestButton("5", "4"),
			e.externalRequestButton("6", "4"),
		),
		container.NewGridWithColumns(5,
			widget.NewLabel("Floor: 3"),
			e.externalRequestButton("1", "3"),
			e.externalRequestButton("2", "3"),
			e.externalRequestButton("5", "3"),
			e.externalRequestButton("6", "3"),
		),
		container.NewGridWithColumns(3,
			widget.NewLabel("Floor: 2"),
			e.externalRequestButton("Up", "2"),
			e.externalRequestButton("Down", "2"),
		),
		container.NewGridWithColumns(6,
			widget.NewLabel("Floor: 1"),
			e.externalRequestButton("2", "1"),
			e.externalRequestButton("3", "1"),
			e.externalRequestButton("4", "1"),
			e.externalRequestButton("5", "1"),
			e.externalRequestButton("6", "1"),
		),
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
