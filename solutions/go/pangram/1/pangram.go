package pangram

import "strings"

func IsPangram(input string) bool {
    input = strings.ToLower(input)
    usedLetters := make(map[rune]bool)
    for _, letter := range input {
        if letter < 97 || letter > 122 {
            continue
        }
        usedLetters[letter] = true
    }
    return len(usedLetters) == 26
}
