package controller

import (
	"github.com/erbilsilik/elevator-navigation-app/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func newElevatorController() *ElevatorController {
	elevator := model.NewElevator(0, "")
	elevatorController := NewElevatorController(
		&[]model.Floor{
			{Name: "P2", IsPressed: false},
			{Name: "P1", IsPressed: false},
			{Name: "1", IsPressed: false},
			{Name: "2", IsPressed: false},
			{Name: "3", IsPressed: false},
			{Name: "4", IsPressed: false},
		},
		elevator,
		time.Second * 1,
		time.Second * 2,
		0,
		nil,
	)
	return elevatorController
}

func Test_GetFloorIndex_WhenFloorExist(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()

	// Act
	floorIndex := elevatorController.getFloorIndex("1")

	// Assert
	assert.Equal(t, 2, floorIndex)
}

func Test_GetFloorIndex_WhenFloorDoes_Not_Exist(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()

	// Act
	floorIndex := elevatorController.getFloorIndex("7")

	// Assert
	assert.Equal(t, -1, floorIndex)
}

func Test_GetFloorFromIndex_WhenFloorExist(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()

	// Act
	floor := elevatorController.getFloorFromIndex(0)

	// Assert
	assert.Equal(t, "P2", floor.Name)
	assert.Equal(t, false, floor.IsPressed)
}

func Test_IsQueueEmpty_WhenIsEmpty(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()

	// Act
	isEmpty := elevatorController.isQueueEmpty()

	// Assert
	assert.Equal(t, true, isEmpty)
}

func Test_IsQueueEmpty_WhenIs_Not_Empty(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 3)

	// Act
	isEmpty := elevatorController.isQueueEmpty()

	// Assert
	assert.Equal(t, false, isEmpty)
}

func Test_IsCurrentFloorIndexLessThanDestinationFloorIndex_WhenIs(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()

	// Act
	isLess := elevatorController.isCurrentFloorIndexLessThanDestinationFloorIndex(3)

	// Assert
	assert.Equal(t, true, isLess)
}

func Test_IsCurrentFloorIndexGreaterThanDestinationFloorIndex_WhenIs(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.currentIndex = 5

	// Act
	isGreater := elevatorController.isCurrentFloorIndexGreaterThanDestinationFloorIndex(3)

	// Assert
	assert.Equal(t, true, isGreater)
}

// <-------------ON PRESS------------->
	// <-------------EXTERNAL REQUESTS------------->
func Test_OnPress_WhenQueueIsEmptyAndRequestIsExternal(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()

	// Act
	elevatorController.OnPress("1", 1)

	// Assert
	assert.Equal(t, 1, len(elevatorController.queue))
}
		// <-------------UP BUTTON PRESSED------------->
func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIsGreaterAndRequestIsExternalAndUpButtonPressed(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 2)
	elevatorController.currentIndex = 3

	// Act
	elevatorController.OnPress("1", 1)

	// Assert
	assert.Equal(t, 2, len(elevatorController.queue))
}

func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIsLessAndRequestIsExternalAndUpButtonPressed(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 4)

	// Act
	elevatorController.OnPress("2", 1)

	// Assert
	isPressed := elevatorController.getFloorFromIndex(3).IsPressed

	assert.Equal(t, true, isPressed)
}

func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIsLessAndPreviousDestinationIsLessAndRequestIsExternalAndUpButtonPressed(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 1)

	// Act
	elevatorController.OnPress("3", 1)

	// Assert
	isPressed := elevatorController.getFloorFromIndex(1).IsPressed
	indexInQueue := elevatorController.queue[0]

	assert.Equal(t, true, isPressed)
	assert.Equal(t, 4, indexInQueue)
}
		// <-------------UP BUTTON PRESSED------------->

		// <-------------DOWN BUTTON PRESSED------------->
func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIsGreaterAndRequestIsExternalAndDownButtonPressed(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 4)
	elevatorController.currentIndex = 4

	// Act
	elevatorController.OnPress("2", -1)
	isPressed := elevatorController.getFloorFromIndex(3).IsPressed

	// Assert
	assert.Equal(t, true, isPressed)
}

func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIsLessAndRequestIsExternalAndDownButtonPressed(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 2)
	elevatorController.currentIndex = 2

	// Act
	elevatorController.OnPress("3", -1)

	// Assert
	assert.Equal(t, 2, len(elevatorController.queue))
}
		// <-------------DOWN BUTTON PRESSED------------->
	// <-------------EXTERNAL REQUESTS------------->

	// <-------------INTERNAL REQUESTS------------->
func Test_OnPress_WhenQueueIsEmptyAndRequestIsInternal(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()

	// Act
	elevatorController.OnPress("1", 0)

	// Assert
	assert.Equal(t, 1, len(elevatorController.queue))
}
		// <-------------IS GOING UP------------->
func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIsLessThanDestinationAndRequestIsInternalAndIsGoingUp(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 5)
	elevatorController.currentIndex = 3
	elevatorController.elevator.Motion = 1

	// Act
	elevatorController.OnPress("1", 0)

	// Assert
	assert.Equal(t, 2, len(elevatorController.queue))
}

func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIs_Not_LessThanDestinationAndRequestIsInternalAndIsGoingUp(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 5)
	elevatorController.currentIndex = 2
	elevatorController.elevator.Motion = 1

	// Act
	elevatorController.OnPress("2", 0)

	// Assert
	isPressed := elevatorController.getFloorFromIndex(3).IsPressed

	assert.Equal(t, true, isPressed)
}
		// <-------------IS GOING DOWN------------->
func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIsGreaterThanDestinationAndRequestIsInternalAndIsGoingDown(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 6)
	elevatorController.currentIndex = 5
	elevatorController.elevator.Motion = -1

	// Act
	elevatorController.OnPress("2", 0)

	// Assert
	isPressed := elevatorController.getFloorFromIndex(3).IsPressed

	assert.Equal(t, true, isPressed)
}

func Test_OnPress_WhenQueueIs_Not_EmptyAndCurrentIndexIs_Not_GreaterThanDestinationAndRequestIsInternalAndIsGoingDown(t *testing.T) {
	// Arrange
	elevatorController := newElevatorController()
	elevatorController.queue = append(elevatorController.queue, 6)
	elevatorController.currentIndex = 2
	elevatorController.elevator.Motion = -1

	// Act
	elevatorController.OnPress("2", 0)

	// Assert
	assert.Equal(t, 2, len(elevatorController.queue))
}
		// <-------------IS GOING DOWN------------->

	// <-------------INTERNAL REQUESTS------------->
// <-------------ON PRESS------------->

// <-------------NAVIGATE------------->