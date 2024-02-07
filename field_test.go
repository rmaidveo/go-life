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
