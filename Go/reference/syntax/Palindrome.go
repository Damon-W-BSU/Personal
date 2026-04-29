package syntax

import (
	"fmt"
	"strings"
)

func IsPalindrome(word string) bool {

	word = strings.ToLower(word)
	letters := strings.Split(word, "")

	for i := range len(word) / 2 {
		if letters[i] != letters[len(word)-1-i] {
			fmt.Printf("\"%s\" is not a palindrome\n", word)
			return false
		}
	}

	fmt.Printf("\"%s\" is a palindrome!\n", word)
	return true
}
