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
