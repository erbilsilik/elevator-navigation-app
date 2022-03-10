package controller

import (
	"fmt"
	"github.com/erbilsilik/elevator-navigation-app/pkg/model"
	"time"
)

type ElevatorController struct {
	floors          *[]model.Floor
	elevator        *model.Elevator
	travelTime      time.Duration
	waitTime        time.Duration
	currentIndex    int
	queue           []int
	firstFloorIndex int
	lastFloorIndex  int
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

func (ec *ElevatorController) setPressedFloor(floorIndex int) {
	(*ec.floors)[floorIndex].IsPressed = true
}

func (ec *ElevatorController) isCurrentFloorIndexLessThanDestinationFloorIndex(destinationFloorIndex int) bool{
	return ec.currentIndex < destinationFloorIndex
}

func (ec *ElevatorController) isCurrentFloorIndexGreaterThanDestinationFloorIndex(destinationFloorIndex int) bool{
	return ec.currentIndex > destinationFloorIndex
}

func (ec *ElevatorController) appendToQueue(destinationFloorIndex int) {
	ec.queue = append(ec.queue, destinationFloorIndex)
}

func (ec *ElevatorController) OnPress(floor string, direction int) {
	request := model.Request{Floor: floor, Direction: direction}

	destinationFloorIndex := ec.getFloorIndex(request.Floor)

	if !ec.isValidDestination(destinationFloorIndex) {
		return
	}

	if !ec.isQueueEmpty() {
		if request.IsExternalRequest() {
			if request.IsUpButtonPressed() {
				if ec.isCurrentFloorIndexLessThanDestinationFloorIndex(destinationFloorIndex) {
					previousDestinationFloorIndex := ec.queue[0]
					if previousDestinationFloorIndex < destinationFloorIndex {
						ec.queue = ec.queue[1:]
						ec.appendToQueue(destinationFloorIndex)
						ec.setPressedFloor(previousDestinationFloorIndex)
					} else {
						ec.setPressedFloor(destinationFloorIndex)
					}
				} else {
					ec.appendToQueue(destinationFloorIndex)
				}
			} else if request.IsDownButtonPressed() {
				if ec.isCurrentFloorIndexGreaterThanDestinationFloorIndex(destinationFloorIndex) {
					ec.setPressedFloor(destinationFloorIndex)
				} else {
					ec.appendToQueue(destinationFloorIndex)
				}
			}
		} else if request.IsInternalRequest() {
			if ec.elevator.IsGoingUp() {
				if ec.isCurrentFloorIndexLessThanDestinationFloorIndex(destinationFloorIndex) {
					ec.setPressedFloor(destinationFloorIndex)
				} else {
					ec.appendToQueue(destinationFloorIndex)
				}
			} else if ec.elevator.IsGoingDown() {
				if ec.isCurrentFloorIndexGreaterThanDestinationFloorIndex(destinationFloorIndex) {
					ec.setPressedFloor(destinationFloorIndex)
				} else {
					ec.appendToQueue(destinationFloorIndex)
				}
			}
		}
	} else {
		ec.appendToQueue(destinationFloorIndex)
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