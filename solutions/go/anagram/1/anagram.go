package anagram

import "strings"

func IsAnagram(word1 string, word2 string) bool {
    word1 = strings.ToLower(word1)
    word2 = strings.ToLower(word2)
    if word1 == word2 || len(word1) != len(word2) {
        return false
    }
    word1Letters := make(map[rune]int)
	word2Letters := make(map[rune]int)
    for _, letter := range word1 {
        word1Letters[letter]++
    }
    for _, letter := range word2 {
        word2Letters[letter]++
    }
    for letter, count := range word1Letters {
        if word2Letters[letter] != count {
            return false
        }
    }
    return true
}

func Detect(subject string, candidates []string) []string {
    var anagrams []string
    for _, word := range candidates {
        if IsAnagram(subject, word) {
            anagrams = append(anagrams, word)
        }
    }
    return anagrams
}
