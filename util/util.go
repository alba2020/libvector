package util

import "math"

// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equals(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Sum(xs []float64) float64 {
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum
}

func Map(xs []float64, f func(float64) float64) []float64 {
	res := make([]float64, len(xs))
	for i, x := range xs {
		res[i] = f(x)
	}
	return res
}

func ToDegrees(radians float64) float64 {
	return 180 * radians / math.Pi
}
