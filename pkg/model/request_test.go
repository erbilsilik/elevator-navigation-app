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
