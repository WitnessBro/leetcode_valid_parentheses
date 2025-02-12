package validParantheses

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var prices = "})"         // Input array
	var expectedResult = true // The expected answer with correct length.
	res := isValid(prices)    // Calls your implementation

	assert.Equal(t, expectedResult, res)
}
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	openB := "([{"
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if strings.Contains(openB, string(s[i])) {
			stack = append(stack, s[i])
		} else {
			switch s[i] {
			case ']':
				{
					if len(stack) != 0 && stack[len(stack)-1] == '[' {
						stack = stack[:len(stack)-1]
						continue
					} else {
						return false
					}
				}
			case ')':
				{
					if len(stack) != 0 && stack[len(stack)-1] == '(' {
						stack = stack[:len(stack)-1]
						continue
					} else {
						return false
					}
				}
			case '}':
				{
					if len(stack) != 0 && stack[len(stack)-1] == '{' {
						stack = stack[:len(stack)-1]
						continue
					} else {
						return false
					}
				}
			}
		}

	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
