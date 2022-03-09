package model

type Request struct {
	Direction 		 int
	Floor      		string
}

func (r *Request) IsExternalRequest() bool  {
	return r.Direction != 0
}

func (r *Request) IsUpButtonPressed() bool  {
	return r.Direction == 1
}

func (r *Request) IsDownButtonPressed() bool  {
	return r.Direction == -1
}
