package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "racecar"
	fmt.Printf("The input string is : %s\n", input)

	isPalindrome := checkPalindrome(input)
	if isPalindrome {
		fmt.Println("The input string is a palindrome")
	} else {
		fmt.Println("The input string is not a palindrome")
	}
}

func checkPalindrome(input string) bool {
	s := strings.ToLower(input)
	s = removeNonAlphanumeric(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func removeNonAlphanumeric(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c <= '0' && c <= '9') {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}
