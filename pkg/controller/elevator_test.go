package controller

import (
	model "github.com/erbilsilik/elevator-navigation-app/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var elevatorController *ElevatorController

func init() {
	// Arrange
	elevator := model.NewElevator(0, "")
	elevatorController = NewElevatorController(
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
		//elevatorView.Display,
	)
}

func Test_GetFloorIndex_WhenFloorExist(t *testing.T) {
	// Act
	floorIndex := elevatorController.getFloorIndex("1")

	// Assert
	assert.Equal(t, 2, floorIndex)
}

func Test_GetFloorIndex_WhenFloorDoes_Not_Exist(t *testing.T) {
	// Act
	floorIndex := elevatorController.getFloorIndex("7")

	// Assert
	assert.Equal(t, -1, floorIndex)
}

func Test_GetFloorFromIndex_WhenFloorExist(t *testing.T) {
	// Act
	floor := elevatorController.getFloorFromIndex(0)

	// Assert
	assert.Equal(t, "P2", floor.Name)
	assert.Equal(t, false, floor.IsPressed)
}

func Test_OnPress_WhenQueueIsEmpty(t *testing.T) {
	// Act
	elevatorController.OnPress("4")

	// Assert
	assert.Equal(t, 5, elevatorController.queue[0])
}

/*
func Test_OnPress_WhenQueueIs_Not_EmptyAndIsLess(t *testing.T) {
	// Arrange
	elevatorController.queue = append(elevatorController.queue, 0)
	elevatorController.currentIndex = 0

	// Act
	elevatorController.OnPress("3")

	// Assert
	assert.Equal(t, 1, len(elevatorController.queue))
}
*/

/*
func Test_OnPress_WhenQueueIs_Not_EmptyAndIsGreater(t *testing.T) {
	// Arrange
	elevatorController.queue = append(elevatorController.queue, 3)
	elevatorController.currentIndex = 4

	// Act
	elevatorController.OnPress("1")

	// Assert
	assert.Equal(t, 2, len(elevatorController.queue))
}
 */

func Test_Handle_WhenArrived(t *testing.T) {
	// Arrange
	elevatorController.queue = append(elevatorController.queue, 4)
	elevatorController.currentIndex = 4

	// Act
	elevatorController.handle()

	// Assert
	assert.Equal(t, 0, len(elevatorController.queue))
}

func Test_Handle_WhenOnFloor(t *testing.T) {
	// Arrange
	elevatorController.queue = append(elevatorController.queue, 1)
	elevatorController.currentIndex = 0
	elevatorController.elevator.Motion = 1

	// Act
	elevatorController.handle()

	// Assert
	assert.Equal(t, 1, elevatorController.currentIndex)
	assert.Equal(t, "P1", elevatorController.elevator.CurrentFloor)
}
