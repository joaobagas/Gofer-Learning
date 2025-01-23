package src

import "fmt"

type Array[T any] struct {
	ColumnCount int
	RowCount    int
	Matrix      [][]T
}

func CreateArray[T any](columnTitles []T) Array[T] {
	var arr Array[T]
	arr.ColumnCount = len(columnTitles)
	arr.RowCount = 1
	arr.Matrix = append(arr.Matrix, columnTitles)
	return arr
}

func AddRow[T any](arr *Array[T], row []T) error {
	if len(row) != arr.ColumnCount {
		return fmt.Errorf("row length (%d) does not match column count (%d)", len(row), arr.ColumnCount)
	}
	arr.RowCount += 1
	arr.Matrix = append(arr.Matrix, row)
	return nil
}

func GetElement[T any](arr *Array[T], row, col int) (T, error) {
	if row < 0 || row >= arr.RowCount || col < 0 || col >= arr.ColumnCount {
		var zero T
		return zero, fmt.Errorf("index out of bounds")
	}
	return arr.Matrix[row][col], nil
}

func Display[T any](arr Array[T]) {
	for _, row := range arr.Matrix {
		fmt.Println(row)
	}
}
