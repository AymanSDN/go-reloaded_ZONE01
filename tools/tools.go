package tools

import (
	"fmt"
	"log"
	"os"
	"unicode"
)

func ContainsLetter(s string) bool {
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			return true
		}
	}
	return false
}

func IsWordApost(s string) bool {
	switch s {
	case "t", "s", "d", "m", "ma", "ll", "ve", "re",
		"T", "S", "D", "M", "MA", "LL", "VE", "RE":

		return true
	default:
		return false
	}
}

func IsVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H':
		return true
	default:
		return false
	}
}

func IsPunct(r rune) bool {
	switch r {
	case '.', ',', ':', ';', '!', '?':
		return true
	default:
		return false
	}
}

func RemoveEmptyString(slice []string) []string {
	var result []string
	for i := 0; i < len(slice); i++ {
		if slice[i] != "" {
			result = append(result, slice[i])
		}
	}
	return result
}

func CheckError(err error, msg string) {
	if err != nil {
		fmt.Fprint(os.Stderr, msg+"\n")
		log.Fatal(err)

	}
}
