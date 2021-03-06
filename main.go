package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/erbilsilik/elevator-navigation-app/pkg/constants"
	"github.com/erbilsilik/elevator-navigation-app/pkg/controller"
	"github.com/erbilsilik/elevator-navigation-app/pkg/model"
	"github.com/erbilsilik/elevator-navigation-app/pkg/view"
	"time"
)

func main() {
	application := app.New()

	elevator := model.NewElevator(0, "")

	elevatorController := controller.NewElevatorController(
		&[]model.Floor{
			{Name: "1", IsPressed: false},
			{Name: "2", IsPressed: false},
			{Name: "3", IsPressed: false},
			{Name: "4", IsPressed: false},
			{Name: "5", IsPressed: false},
			{Name: "6", IsPressed: false},
		},
		elevator,
		time.Second * constants.TravelTime,
		time.Second * constants.WaitTime,
		0,
		nil,
	)

	elevatorView := view.NewElevator(elevatorController)
	elevatorView.LoadUI(application)

	application.Run()
}
