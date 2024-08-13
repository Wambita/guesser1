package functionFiles

// calculates the mean of a  slice of float numbers
func Mean(numslice []float64) float64 {
	sum := 0.0
	sliceLength := len(numslice)
	for _, num := range numslice {
		sum += num
	}
	mean := sum / float64(sliceLength)
	return mean
}
