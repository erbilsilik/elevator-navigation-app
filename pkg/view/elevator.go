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
		e.elevatorController.OnPress(floor, 0)
	})
}

func (e *Elevator) externalRequestButton(floorDirection string, sourceFloor string) *widget.Button {
	return e.addButton(floorDirection, func() {
		var direction int
		if floorDirection == "Up" {
			direction = 1
		} else if floorDirection == "Down"{
			direction = -1
		} else {
			direction = 0
		}
		e.elevatorController.OnPress(sourceFloor, direction)
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
		container.NewGridWithColumns(2,
			widget.NewLabel("Floor: 6"),
			e.externalRequestButton("Down", "6"),
		),
		container.NewGridWithColumns(3,
			widget.NewLabel("Floor: 5"),
			e.externalRequestButton("Up", "5"),
			e.externalRequestButton("Down", "5"),
		),
		container.NewGridWithColumns(3,
			widget.NewLabel("Floor: 4"),
			e.externalRequestButton("Up", "4"),
			e.externalRequestButton("Down", "4"),
		),
		container.NewGridWithColumns(3,
			widget.NewLabel("Floor: 3"),
			e.externalRequestButton("Up", "3"),
			e.externalRequestButton("Down", "3"),
		),
		container.NewGridWithColumns(3,
			widget.NewLabel("Floor: 2"),
			e.externalRequestButton("Up", "2"),
			e.externalRequestButton("Down", "2"),
		),
		container.NewGridWithColumns(2,
			widget.NewLabel("Floor: 1"),
			e.externalRequestButton("Up", "1"),
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
