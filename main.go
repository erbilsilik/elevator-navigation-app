package main

import (
	"fyne.io/fyne/v2/app"
	controller2 "github.com/erbilsilik/elevator-navigation-app/pkg/controller"
	model2 "github.com/erbilsilik/elevator-navigation-app/pkg/model"
	view2 "github.com/erbilsilik/elevator-navigation-app/pkg/view"
	"time"
)

func main() {
	application := app.New()

	elevator := model2.NewElevator(0, "")

	elevatorController := controller2.NewElevatorController(
		&[]model2.Floor{
			{Name: "P2", IsPressed: false},
			{Name: "P1", IsPressed: false},
			{Name: "1", IsPressed: false},
			{Name: "2", IsPressed: false},
			{Name: "3", IsPressed: false},
			{Name: "4", IsPressed: false},
		},
		elevator,
		time.Second * 5,
		time.Second * 10,
		0,
		nil,
	)

	elevatorView := view2.NewElevator(elevatorController)
	elevatorView.LoadUI(application)

	application.Run()
}
