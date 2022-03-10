package model

type Elevator struct {
	Motion       int
	CurrentFloor string
}

func (e Elevator) IsGoingUp() bool {
	return e.Motion == 1
}

func (e Elevator) IsGoingDown() bool {
	return e.Motion == -1
}

func (e Elevator) IsIdle() bool {
	return e.Motion == 0
}

func NewElevator(motion int, currentFloor string) *Elevator {
	return &Elevator{
		Motion: motion,
		CurrentFloor: currentFloor,
	}
}