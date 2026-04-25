package acronym

import "strings"

// Abbreviate should have a comment documenting it.
func Abbreviate(text string) string {
    text = strings.ToUpper(text)
    replacer := strings.NewReplacer("-", " ", "_", " ")
	text = replacer.Replace(text)
    var words []string = strings.Fields(text)
    var abbreviationLetters []byte
    for _, word := range words {
        abbreviationLetters = append(abbreviationLetters, word[0])
    }
	return string(abbreviationLetters)
}
