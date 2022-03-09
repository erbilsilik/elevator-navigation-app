package controller

import (
	"fmt"
	"github.com/erbilsilik/elevator-navigation-app/pkg/model"
	"time"
)

type ElevatorController struct {
	floors       *[]model.Floor
	elevator     *model.Elevator
	travelTime   time.Duration
	waitTime   	 time.Duration
	currentIndex int
	queue        []int
}

func (ec *ElevatorController) getFloorIndex(element string) int {
	for k, v := range *ec.floors {
		if element == v.Name {
			return k
		}
	}
	return -1
}

func (ec *ElevatorController) getFloorFromIndex(index int) model.Floor {
	for k, v := range *ec.floors {
		if k == index {
			return model.Floor{Name: v.Name, IsPressed: v.IsPressed}
		}
	}
	return model.Floor{}
}

func (ec *ElevatorController) isValidDestination(destinationFloorIndex int) bool {
	if destinationFloorIndex < 0 {
		return false
	}
	return true
}

func (ec *ElevatorController) isQueueEmpty() bool {
	return len(ec.queue) == 0
}

func (ec *ElevatorController) isElevatorGoingUp() bool {
	return ec.elevator.Motion == 1
}

func (ec *ElevatorController) isElevatorGoingDown() bool {
	return ec.elevator.Motion == -1
}

func (ec *ElevatorController) isCurrentFloorLessThanDestinationFloor(destinationFloorIndex int) bool {
	return ec.currentIndex < destinationFloorIndex
}

func (ec *ElevatorController) isCurrentFloorGreaterThanDestinationFloor(destinationFloorIndex int) bool {
	return ec.currentIndex > destinationFloorIndex
}

func (ec *ElevatorController) updateDestinations(destinationFloorIndex int, lastDestinationFloorIndex int) {
	if destinationFloorIndex < lastDestinationFloorIndex {
		(*ec.floors)[destinationFloorIndex].IsPressed = true
	} else {
		(*ec.floors)[lastDestinationFloorIndex].IsPressed = true
		ec.queue = ec.queue[1:]
		ec.queue = append(ec.queue, destinationFloorIndex)
	}
}

func (ec *ElevatorController) OnPress(sourceFloor string, direction int) {
	request := model.Request{SourceFloor: sourceFloor, Direction: direction}

	destinationFloorIndex := ec.getFloorIndex(request.SourceFloor)
	if !ec.isValidDestination(destinationFloorIndex) {
		return
	}
	if !ec.isQueueEmpty() {
		lastDestinationFloorIndex := ec.queue[0]
		if ec.isCurrentFloorLessThanDestinationFloor(destinationFloorIndex) && ec.isElevatorGoingUp() ||
			ec.isCurrentFloorGreaterThanDestinationFloor(destinationFloorIndex) && ec.isElevatorGoingDown() {
				if request.IsExternalRequest() {
					var florIndex int
					if request.ShouldGoUp() {
						florIndex = ec.getFloorIndex("6")
						ec.updateDestinations(florIndex, lastDestinationFloorIndex)
						(*ec.floors)[destinationFloorIndex].IsPressed = true
					} else if request.ShouldGoDown() {
						florIndex = ec.getFloorIndex("1")
						ec.updateDestinations(florIndex, lastDestinationFloorIndex)
						(*ec.floors)[destinationFloorIndex].IsPressed = true
					}
				} else {
					ec.updateDestinations(destinationFloorIndex, lastDestinationFloorIndex)
				}
		} else {
			ec.queue = append(ec.queue, destinationFloorIndex)
		}
	} else {
		ec.queue = append(ec.queue, destinationFloorIndex)
		if request.IsExternalRequest() {
			var nextFloorIndex int
			if request.ShouldGoDown() {
				nextFloorIndex = ec.getFloorIndex("1")
			} else if request.ShouldGoUp() {
				nextFloorIndex = ec.getFloorIndex("6")
			}
			ec.queue = append(ec.queue, nextFloorIndex)
		}
		go ec.navigate()
	}
}

func (ec *ElevatorController) navigate() {
	if len(ec.queue) == 0 {
		return
	}
	if ec.currentIndex == ec.queue[0] {
		ec.elevator.Motion = 0
		ec.queue = ec.queue[1:]
		ec.fireEvent("arrived")
		if len(ec.queue) == 0 {
			return
		}
	}
	if ec.elevator.Motion != 0 {
		ec.currentIndex += ec.elevator.Motion
		ec.elevator.CurrentFloor = ec.getFloorFromIndex(ec.currentIndex).Name
		ec.fireEvent("floor")
		if ec.getFloorFromIndex(ec.currentIndex).IsPressed {
			fmt.Print("waiting on floor...")
			time.Sleep(ec.waitTime)
		}
	}
	if len(ec.queue) > 0 {
		if ec.currentIndex < ec.queue[0] {
			ec.elevator.Motion = +1
			fmt.Print("going up...")
			time.Sleep(ec.travelTime)
			ec.fireEvent("up")
		} else if ec.currentIndex > ec.queue[0] {
			fmt.Print("going down...")
			ec.elevator.Motion = -1
			time.Sleep(ec.travelTime)
			ec.fireEvent("down")
		}
	}
	(*ec.floors)[ec.currentIndex].IsPressed = false
	ec.navigate()
}

func (ec *ElevatorController) fireEvent(floor string) {
	fmt.Println(floor + " : " + ec.elevator.CurrentFloor)
}

func NewElevatorController(floors *[]model.Floor, elevator *model.Elevator, travelTime time.Duration, waitTime time.Duration,
	currentIndex int, queue []int) *ElevatorController {
	return &ElevatorController{
		floors: floors,
		elevator: elevator,
		travelTime: travelTime,
		waitTime: waitTime,
		currentIndex: currentIndex,
		queue: queue,
	}
}