package functionFiles

import "math"

func Stddev(variance float64) float64 {
	return math.Sqrt(variance)
}

func Variance(nums []float64) float64 {
	lenNum := len(nums)
	var sum float64

	for _, val := range nums {
		sum += math.Pow(val-Mean(nums), 2.0)
	}
	return sum / float64(lenNum)
}
