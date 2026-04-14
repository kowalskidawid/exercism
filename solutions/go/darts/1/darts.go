package darts

import "math"

func Score(x, y float64) int {
    z := math.Sqrt(x * x + y * y)
    switch {
        case z <= 1:
        	return 10
        case z <= 5:
        	return 5
        case z <= 10:
        	return 1
    }
    return 0
}
