package scrabblescore

import "strings"

func Score(word string) int {
    score := 0
    word = strings.ToUpper(word)
    for _, letter := range word {
        switch {
            case letter == 'D' || letter == 'G':
            	score += 2
            case letter == 'B' || letter == 'C' || letter == 'M' || letter == 'P':
            	score += 3
            case letter == 'F' || letter == 'H' || letter == 'V' || letter == 'W' || letter == 'Y':
            	score += 4
            case letter == 'K':
            	score += 5
            case letter == 'J' || letter == 'X':
            	score += 8
            case letter == 'Q' || letter == 'Z':
            	score += 10
            default:
                score += 1
        }
    }
	return score
}
