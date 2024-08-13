package functionFiles

import "math"

// calculates standard deviation from variance
func Stddev(variance float64) float64 {
	return math.Sqrt(variance)
}

// calculates variance from mean
func Variance(nums []float64) float64 {
	lenNum := len(nums)
	var sum float64

	for _, val := range nums {
		sum += math.Pow(val-Mean(nums), 2.0)
	}
	return sum / float64(lenNum)
}
