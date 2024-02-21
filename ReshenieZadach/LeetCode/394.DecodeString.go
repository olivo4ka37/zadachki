package LeetCode

import (
	"strconv"
	"strings"
)

func DecodeString(s string) string {
	stack := []string{}     // Initialize a stack to store characters and substrings
	var res strings.Builder // Use a strings.Builder for efficient string concatenation

	for _, c := range s {
		if c != ']' { // If the character is not ']', push it onto the stack as a string
			stack = append(stack, string(c))
		} else {
			substr := ""                     // Initialize an empty string to store the substring
			for stack[len(stack)-1] != "[" { // Pop characters from the stack until '[' is found
				substr = stack[len(stack)-1] + substr
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1] // Pop the '[' from the stack

			k := ""
			for len(stack) > 0 && stack[len(stack)-1] >= "0" && stack[len(stack)-1] <= "9" {
				k = stack[len(stack)-1] + k
				stack = stack[:len(stack)-1]
			}

			num, _ := strconv.Atoi(k) // Parse the repeat count as an integer

			repeated := strings.Repeat(substr, num) // Create the repeated substring
			stack = append(stack, repeated)         // Push the repeated substring onto the stack
		}
	}

	// Build the result string using the strings.Builder
	for _, s := range stack {
		res.WriteString(s)
	}

	return res.String() // Return the decoded string as a string
}
