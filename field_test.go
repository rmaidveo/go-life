package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewField(test *testing.T) {
	actualResult := NewField(3, 2)

	expectedResult := Field{
		{false, false, false},
		{false, false, false},
	}
	assert.Equal(test, expectedResult, actualResult)
}

func TestField_Width(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	result := field.Width()

	assert.Equal(test, 3, result)
}

func TestField_Heigth(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	result := field.Height()

	assert.Equal(test, 2, result)
}

func TestField_Cell(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, true},
	}
	result := field.Cell(2, 1)

	assert.True(test, result)
}

func TestField_Cell_withCoordinatesBeyondMinimum(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, true, false},
	}
	result := field.Cell(-2, -1)

	assert.True(test, result)
}

func TestField_Cell_withCoordinatesBeyondMaximum(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, true, false},
	}
	result := field.Cell(4, 3)

	assert.True(test, result)
}

func TestField_SetCell(test *testing.T) {
	field := Field{
		{false, false, false},
		{false, false, false},
	}
	field.SetCell(2, 1, true)

	expectedField := Field{
		{false, false, false},
		{false, false, true},
	}
	assert.Equal(test, expectedField, field)
}

func TestField_NeighborCount(test *testing.T) {
	field := Field{
		{true, false, false},
		{false, true, true},
		{true, false, false},
	}
	result := field.NeighborCount(0, 1)

	assert.Equal(test, 4, result)
}

func TestField_NeighborCount_forAliveCell(test *testing.T) {
	field := Field{
		{true, false, false},
		{true, true, true},
		{true, false, false},
	}
	result := field.NeighborCount(0, 1)

	assert.Equal(test, 4, result)
}

func TestField_NextCell_willBeBorn(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false /* ! */, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	result := field.NextCell(1, 2)

	assert.True(test, result)
}

func TestField_NextCell_willSurvive(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true /* ! */, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	result := field.NextCell(3, 2)

	assert.True(test, result)
}

func TestField_NextCell_willDie(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true /* ! */, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	result := field.NextCell(2, 1)

	assert.False(test, result)
}

func TestField_NextField(test *testing.T) {
	field := Field{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, false, true, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
	}
	actualField := field.NextField()

	expectedNextField := Field{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, false, true, false},
		{false, false, true, true, false},
		{false, false, true, false, false},
	}
	assert.Equal(test, expectedNextField, actualField)
}

func TestField_String(test *testing.T) {
	field := Field{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	result := field.String()

	assert.Equal(test, ".O.\n..O\nOOO\n", result)
}

func TestParseField_succses(test *testing.T) {
	text := "!comment #1\n!comment #2\n.O.\n\n..O\nOOO"
	actualResult, err := ParseField(text)

	expectedResult := Field{
		{false, true, false},
		{false, false, true},
		{true, true, true},
	}
	assert.Equal(test, expectedResult, actualResult)
	assert.NoError(test, err)
}

func TestParseField_error_unknownCharacter(test *testing.T) {
	result, err := ParseField(".O.\n..*\nOOO")

	assert.Nil(test, result)
	assert.EqualError(test, err, "unknown character '*'")
}

func TestParseField_error_inconsistentLengthOfLine(test *testing.T) {
	result, err := ParseField(".O.\n..\nOOO")

	assert.Nil(test, result)
	assert.EqualError(test, err, "inconsistent length of line 1")
}
