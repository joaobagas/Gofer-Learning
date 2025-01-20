package src

// y = alpha * x + beta
type Line struct {
	Alpha float32
	Beta  float32
}

func CreateLinearRegression() Line {
	newLine := Line{
		Alpha: 1,
		Beta:  0,
	}
	return newLine
}

func FitLinearRegression(line *Line, x []float32, y []float32) {
	if len(x) != len(y) {
		return
	}

	sumX, sumY, sumXY, sumXSQ := float32(0), float32(0), float32(0), float32(0)
	for i := 0; i < len(x); i++ {
		sumX += x[i]
		sumY += y[i]
		sumXY += x[i] * y[i]
		sumXSQ += x[i] * x[i]
	}

	n := float32(len(x))
	denominator := (n * sumXSQ) - (sumX * sumX)
	if denominator == 0 {
		return
	}

	line.Alpha = ((n * sumXY) - (sumX * sumY)) / denominator
	line.Beta = (sumY - line.Alpha*sumX) / n
}

func PredictLinearRegression(line Line, val float32) float32 {
	return val*line.Alpha + line.Beta
}
