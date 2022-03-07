package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewElevator_WhenDefault(t *testing.T) {
	// arrange
	elevator := NewElevator(0, "")

	// act

	// assert
	assert.Equal(t, "", elevator.CurrentFloor)
	assert.Equal(t, 0, elevator.Motion)
}
