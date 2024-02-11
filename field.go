package life

import "strings"

const (
	deadCellRepresentation  = '.'
	aliveCellRepresentation = 'O'
)

type Field [][]bool

func NewField(width int, hieght int) Field {
	field := make(Field, hieght)
	for index := range field {
		field[index] = make([]bool, width)
	}

	return field
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
