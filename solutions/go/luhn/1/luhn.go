package luhn

import (
    "strings"
    "slices"
    "strconv"
)
func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    var sum int
    digits := []rune(id)
    if len(digits) <= 1 {
        return false
    }
    slices.Reverse(digits)
    for index, numberSign := range digits {
        number, error := strconv.Atoi(string(numberSign))
        if error != nil {
            return false
        }
        if index % 2 != 0 {
           	number *= 2
            if number > 9 {
                number -= 9
            }
        }
        sum += number
    }
    return sum % 10 == 0
}
