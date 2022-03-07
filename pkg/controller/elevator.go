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

func (ec *ElevatorController) isElevatorGoingUp(floorIndex int) bool {
	return ec.currentIndex < floorIndex && ec.elevator.Motion == 1
}

func (ec *ElevatorController) isElevatorGoingDown(floorIndex int) bool {
	return ec.currentIndex > floorIndex && ec.elevator.Motion == -1
}

func (ec *ElevatorController) OnPress(floor string) {
	index := ec.getFloorIndex(floor)
	if index < 0 {
		return
	}
	if len(ec.queue) != 0 {
		floorIndex := ec.getFloorIndex(floor)
		if ec.isElevatorGoingUp(floorIndex) || ec.isElevatorGoingDown(floorIndex) {
			(*ec.floors)[floorIndex].IsPressed = true
		} else {
			ec.queue = append(ec.queue, index)
		}
	} else {
		ec.queue = append(ec.queue, index)
		go ec.handle()
	}
}

func (ec *ElevatorController) handle() {
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
	ec.handle()
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