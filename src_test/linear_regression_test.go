package src_test

import (
	"gofer-learning/src"
	"testing"
)

func TestRegession(t *testing.T) {
	x := []float32{1, 2, 3, 4, 5}
	y := []float32{2, 4, 5, 4, 5}

	line := src.CreateLinearRegression()
	src.FitLinearRegression(&line, x, y)

	val := float32(6)
	prediction := src.PredictLinearRegression(line, val)

	if line.Alpha != 0.6 {
		t.Fatalf("Alpha wasn't calculated properly.")
	}
	if line.Beta != 2.2 {
		t.Fatalf("Beta wasn't calculated properly.")
	}
	if prediction != 5.8 {
		t.Fatalf("Prediction wasn't calculated properly.")
	}
}
