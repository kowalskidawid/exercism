package reversestring

import "slices"

func Reverse(input string) string {
    letters := []rune(input)
    slices.Reverse(letters)
	return string(letters)
}
