package life

type Field [][]bool

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

func wrapAroundModulus(number int, modulus int) int {
	number += modulus
	return number % modulus
}
