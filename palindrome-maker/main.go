package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "raceking"
	fmt.Printf("The original word is: %s\n", input)

	palindrome := makePalindrome(input)
	fmt.Printf("ThePalindrome of %s is: %s\n", input, palindrome)
}

func makePalindrome(input string) string {
	var sb strings.Builder
	sb.WriteString(input)

	for i := len(input) - 1; i >= 0; i-- {
		sb.WriteByte(input[i])
	}

	return sb.String()
}
