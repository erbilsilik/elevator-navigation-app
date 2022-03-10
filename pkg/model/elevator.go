package model

import "github.com/erbilsilik/elevator-navigation-app/pkg/constants"

type Elevator struct {
	Direction    constants.Direction
	CurrentFloor string
}

func (e *Elevator) IsGoingUp() bool {
	return e.Direction == constants.Up
}

func (e *Elevator) IsGoingDown() bool {
	return e.Direction == constants.Down
}

func (e *Elevator) IsIdle() bool {
	return e.Direction == constants.Idle
}

func (e *Elevator) IsMoving() bool {
	return e.Direction != constants.Idle
}

func (e *Elevator) Compare(currentFloorIndex int, destinationFloorIndex int) bool {
	if e.IsGoingUp() {
		return currentFloorIndex < destinationFloorIndex
	}
	if e.IsGoingDown() {
		return currentFloorIndex > destinationFloorIndex
	}
	return false
}

func NewElevator(direction constants.Direction, currentFloor string) *Elevator {
	return &Elevator{
		Direction: direction,
		CurrentFloor: currentFloor,
	}
}