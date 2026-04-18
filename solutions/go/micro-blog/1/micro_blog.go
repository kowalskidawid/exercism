package microblog

func Truncate(phrase string) string {
	letters := []rune(phrase)
    if len(letters) < 5 {
        return phrase
    }
	return string(letters[:5])
}
