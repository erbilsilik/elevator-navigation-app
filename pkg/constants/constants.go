package constants

type Direction int

const(
	Idle	Direction = 0
	Up      Direction = 1
	Down 	Direction = -1
)

func (d Direction) Int() int{
	return int(d)
}

const TravelTime = 5
const WaitTime = 10

