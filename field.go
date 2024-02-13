package life

import (
	"fmt"
	"strings"
)

const (
	deadCellRepresentation  = '.'
	aliveCellRepresentation = 'O'
)

type Field [][]bool

func NewField(width int, height int) Field {
	field := make(Field, height)
	for index := range field {
		field[index] = make([]bool, width)
	}

	return field
}

func ParseField(text string) (Field, error) {
	var field Field
	lines := strings.Split(text, "\n")
	width := -1
	for index, line := range lines {
		if line == "" || line[0] == '!' {
			continue
		}

		if width == -1 {
			width = len(line)
		} else if len(line) != width {
			return nil, fmt.Errorf("inconsistent length of line %d", index)
		}

		var row []bool
		for _, character := range line {
			switch character {
			case deadCellRepresentation:
				row = append(row, false)
			case aliveCellRepresentation:
				row = append(row, true)
			default:
				return nil, fmt.Errorf("unknown character %q", character)
			}
		}

		field = append(field, row)
	}

	return field, nil
}

func (field Field) Width() int {
	return len(field[0])
}

func (field Field) Height() int {
	return len(field)
}

func (field Field) Cell(column int, row int) bool {
	column = wrapAroundModulus(column, field.Width())
	row = wrapAroundModulus(row, field.Height())
	return field[row][column]
}

func (field Field) SetCell(column int, row int, value bool) {
	field[row][column] = value
}

func (field Field) NeighborCount(column int, row int) int {
	var count int
	for y := row - 1; y <= row+1; y++ {
		for x := column - 1; x <= column+1; x++ {
			if field.Cell(x, y) {
				count++
			}
		}
	}

	if field.Cell(column, row) {
		count--
	}

	return count
}

func (field Field) NextCell(column int, row int) bool {
	cell := field.Cell(column, row)
	neighborCount := field.NeighborCount(column, row)
	willBeBorn := !cell && neighborCount == 3
	willSurvive := cell && (neighborCount == 2 || neighborCount == 3)
	return willBeBorn || willSurvive
}

func (field Field) NextField() Field {
	nextField := NewField(field.Width(), field.Height())
	for y := 0; y < field.Height(); y++ {
		for x := 0; x < field.Width(); x++ {
			nextCell := field.NextCell(x, y)
			nextField.SetCell(x, y, nextCell)
		}
	}

	return nextField
}

func (field Field) String() string {
	var result strings.Builder
	for y := 0; y < field.Height(); y++ {
		for x := 0; x < field.Width(); x++ {
			if field.Cell(x, y) {
				result.WriteRune(aliveCellRepresentation)
			} else {
				result.WriteRune(deadCellRepresentation)
			}
		}

		result.WriteRune('\n')
	}

	return result.String()
}

func wrapAroundModulus(number int, modulus int) int {
	number += modulus
	return number % modulus
}
