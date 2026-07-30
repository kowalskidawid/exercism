package etl

import (
    "strings"
)

func Transform(in map[int][]string) map[string]int {
    var output map[string]int = make(map[string]int)
    for points, letters := range in {
        for _, letter := range letters {
            lowwerCaseLetter := strings.ToLower(letter)
            output[lowwerCaseLetter] = points
        }
    }
    return output
}
