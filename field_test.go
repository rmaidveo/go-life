package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestField_Width(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	result := field.Width()

	assert.Equal(test, 3, result)
}

func TestField_Heigth(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	result := field.Height()

	assert.Equal(test, 2, result)
}

func TestField_Cell(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, true},
	}
	result := field.Cell(2, 1)

	assert.True(test, result)
}

func TestField_Cell_withCoordinatesBeyondMinimum(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, true, false},
	}
	result := field.Cell(-2, -1)

	assert.True(test, result)
}

func TestField_Cell_withCoordinatesBeyondMaximum(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, true, false},
	}
	result := field.Cell(4, 3)

	assert.True(test, result)
}

func TestField_SetCell(test *testing.T) {
	field := Field{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}
	field.SetCell(2, 1, true)

	expectedField := Field{
		[]bool{false, false, false},
		[]bool{false, false, true},
	}
	assert.Equal(test, expectedField, field)
}
