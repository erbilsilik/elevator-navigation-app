package model

import "github.com/erbilsilik/elevator-navigation-app/pkg/constants"

type Request struct {
	Direction constants.Direction
	Floor     string
}

func (r *Request) IsExternalRequest() bool  {
	return r.Direction != constants.Idle
}

func (r *Request) IsInternalRequest() bool  {
	return r.Direction == constants.Idle
}

func (r *Request) IsUpButtonPressed() bool  {
	return r.Direction == constants.Up
}

func (r *Request) IsDownButtonPressed() bool  {
	return r.Direction == constants.Down
}

func (r *Request) Compare(currentIndex int, destinationFloorIndex int) bool {
	if r.IsUpButtonPressed() {
		return currentIndex < destinationFloorIndex
	}
	if r.IsDownButtonPressed() {
		return currentIndex > destinationFloorIndex
	}
	return false
}

func (r *Request) GetIndexesByButtonPress(index1 int, index2 int) (int, int) {
	if r.IsUpButtonPressed() {
		return index1, index2
	}
	if r.IsDownButtonPressed() {
		return index2, index1
	}

	return -1, -1
}
