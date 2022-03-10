package model

import (
	"github.com/erbilsilik/elevator-navigation-app/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
)

// <-------------EXTERNAL REQUESTS------------->
func Test_IsExternalRequest_WhenDirectionIsUp(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Up}

	// Act
	isExternalRequest := request.IsExternalRequest()

	// Assert
	assert.Equal(t, true, isExternalRequest)
}

func Test_IsExternalRequest_WhenDirectionIsDown(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Down}

	// Act
	isExternalRequest := request.IsExternalRequest()

	// Assert
	assert.Equal(t, true, isExternalRequest)
}

func Test_IsExternalRequest_WhenUpButtonPressed(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Up}

	// Act
	isUpButtonPressed := request.IsUpButtonPressed()

	// Assert
	assert.Equal(t, true, isUpButtonPressed)
}

func Test_IsExternalRequest_WhenDownButtonPressed(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Down}

	// Act
	isDownButtonPressed := request.IsDownButtonPressed()

	// Assert
	assert.Equal(t, true, isDownButtonPressed)
}
// <-------------EXTERNAL REQUESTS------------->


// <-------------INTERNAL REQUESTS------------->
func Test_IsInternalRequest_WhenDirectionIsDown(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Idle}

	// Act
	isInternalRequest := request.IsInternalRequest()

	// Assert
	assert.Equal(t, true, isInternalRequest)
}
// <-------------INTERNAL REQUESTS------------->


func Test_Compare_WhenUpButtonPressed(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Up}

	// Act
	result := request.Compare(1, 5)

	// Assert
	assert.Equal(t, true, result)
}

func Test_Compare_WhenDownButtonPressed(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Down}

	// Act
	result := request.Compare(5, 1)

	// Assert
	assert.Equal(t, true, result)
}

func Test_GetIndexesByButtonPress_WhenUpButtonPressed(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Up}
	first := 1
	second := 5

	// Act
	i1, i2 := request.GetIndexesByButtonPress(first, second)

	// Assert
	assert.Equal(t, first, i1)
	assert.Equal(t, second, i2)
}

func Test_GetIndexesByButtonPress_WhenDownButtonPressed(t *testing.T) {
	// Arrange
	request := Request{Floor: "3", Direction: constants.Down}
	first := 5
	second := 1

	// Act
	i1, i2 := request.GetIndexesByButtonPress(first, second)

	// Assert
	assert.Equal(t, first, i2)
	assert.Equal(t, second, i1)
}