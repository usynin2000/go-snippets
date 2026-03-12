package testing_utils

import (
	"math"
)

// Abs returns absolute value.
// For example: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
func Abs(value float64) float64 {
	return math.Abs(value)
}
