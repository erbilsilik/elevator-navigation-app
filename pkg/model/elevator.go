package model

type Elevator struct {
	Motion       int
	CurrentFloor string
}

func NewElevator(motion int, currentFloor string) *Elevator {
	return &Elevator{
		Motion: motion,
		CurrentFloor: currentFloor,
	}
}