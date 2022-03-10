package model

import (
	"github.com/erbilsilik/elevator-navigation-app/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_IsGoingUp_WhenIs(t *testing.T) {
	// Arrange
	elevator := NewElevator(constants.Up, "3")

	// Act
	isGoingUp := elevator.IsGoingUp()

	// Assert
	assert.Equal(t, true, isGoingUp)
}

func Test_IsGoingDown_WhenIs(t *testing.T) {
	// Arrange
	elevator := NewElevator(constants.Down, "3")

	// Act
	isGoingDown := elevator.IsGoingDown()

	// Assert
	assert.Equal(t, true, isGoingDown)
}

func Test_IsIdle_WhenIs(t *testing.T) {
	// Arrange
	elevator := NewElevator(constants.Idle, "3")

	// Act
	isIdle := elevator.IsIdle()

	// Assert
	assert.Equal(t, true, isIdle)
}

func Test_IsMoving_WhenIs(t *testing.T) {
	// Arrange
	elevator := NewElevator(constants.Up, "3")

	// Act
	isMoving := elevator.IsMoving()

	// Assert
	assert.Equal(t, true, isMoving)
}
