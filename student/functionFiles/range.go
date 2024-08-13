package functionFiles

import "math"

// Range calculates a smaller range by using smaller multipliers for standard deviation.
func Range(numslice []float64) (int, int) {
	// Determine the starting index for the subset of the slice to consider
	startCalc := len(numslice) - 4
	if startCalc < 0 {
		startCalc = 0
	}

	// Get the subset of the slice for calculation
	chunk := numslice[startCalc:]
	mean := Mean(chunk)
	variance := Variance(chunk)
	stddev := Stddev(variance)

	// calculating the range
	// use co-efficient to determine the spread of the range
	lowerLimit := mean - (1.8 * stddev)
	upperLimit := mean + (1.8 * stddev)
	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}
