package model

type Request struct {
	Direction 		 int
	SourceFloor      string
}

func (r *Request) IsExternalRequest() bool  {
	return r.Direction != 0
}

func (r *Request) ShouldGoUp() bool  {
	return r.Direction == 1
}

func (r *Request) ShouldGoDown() bool  {
	return r.Direction == -1
}
