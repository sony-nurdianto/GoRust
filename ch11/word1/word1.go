package word1

import "unicode"

func IsPalindrome(s string) bool {
	leters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			leters = append(leters, unicode.ToLower(r))
		}
	}
	for i := range leters {
		if leters[i] != leters[len(leters)-1-i] {
			return false
		}
	}
	return true
}
