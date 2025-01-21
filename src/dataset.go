package src

type Array struct {
	ColumnCount int
	RowCount    int
	Columns     []string
	Rows        []int
}

// Here we define the rows in advance
func CreateArray(columns []string) Array {
	var arr Array

	arr.Columns = columns
	arr.ColumnCount = len(columns)

	return arr
}

func AddRow(arr *Array, row []int) {
	arr.RowCount += 1
	arr.Rows = append(arr.Rows, row)
}

func ToString() {

}
