package isogram

import "unicode"

func IsIsogram(word string) bool {
	uniqueLetters := make(map[rune]bool)
    for _, letter := range word {
        if letter == ' ' || letter == '-' {
            continue
        }
        smallLetter := unicode.ToLower(letter)
        if uniqueLetters[smallLetter] {
            return false
        }
        uniqueLetters[smallLetter] = true
    }
    return true
}
