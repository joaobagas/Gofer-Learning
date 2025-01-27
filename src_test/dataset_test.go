package src_test

import (
	"gofer-learning/src"
	"testing"
)

func TestDataset(t *testing.T) {
	arr := src.CreateArray([]string{"Name", "Age", "City"})
	_ = src.AddRow(&arr, []string{"Alice", "30", "New York"})

	name, _ := src.GetElement(&arr, 1, 0)
	age, _ := src.GetElement(&arr, 1, 1)
	city, _ := src.GetElement(&arr, 1, 2)

	if name != "Alice" || age != "30" || city != "New York" {
		t.Fatalf("Information is wrong!")
	}
}
