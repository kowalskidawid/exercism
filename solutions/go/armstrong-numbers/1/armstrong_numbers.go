package armstrongnumbers

import (
    "math"
    "strconv"
)

func IsNumber(n int) bool {
	nText := strconv.Itoa(n)
    sum := 0
    for _, digitText := range nText {
        digit := int(digitText - '0')
        sum += int(math.Pow(float64(digit), float64(len(nText))))
    }
    return sum == n
}
