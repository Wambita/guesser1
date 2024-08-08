package functionFiles

func Mean(numslice []float64) float64 {
	sum := 0.0
	lenNum := len(numslice)
	for _, num := range numslice {
		sum += num
	}
	mean := sum / float64(lenNum)
	return mean
}
