package model

type Request struct {
	DestinationFloor string
	SourceFloor      string
}

func (r *Request) IsExternalRequest() bool  {
	return r.SourceFloor != ""
}
