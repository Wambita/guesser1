package functionFiles

import "math"

func Range(numslice []float64) (int, int) {
	// groups
	startCalc := len(numslice) - 4
	if startCalc < 0 {
		startCalc = 0
	}
	// chunk to store the elements in chunks
	chunk := numslice[startCalc:]
	mean := Mean(chunk)
	variance := Variance(chunk)
	stddev := Stddev(variance)

	// calculating the range
	lowerLimit := mean - (1.8 * stddev)
	upperLimit := mean + (1.8 * stddev)
	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}
