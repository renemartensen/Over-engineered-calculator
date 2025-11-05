package calculator

import (
	"fmt"
	"regexp"
	"strings"
)

func Tokenize(input string) ([]string, error) {
	input = strings.ReplaceAll(input, " ", "")

	// Define a regex for tokens:
	// \d+\.?\d*   -> integer or decimal number
	// [()+\-*/]   -> operators and parentheses
	re := regexp.MustCompile(`\d+\.?\d*|[()+\-*/]`)

	tokens := re.FindAllString(input, -1)

	// lets say there is an invalid character, then the regex would skip it
	// we can check this by reconstructing the input from the tokens and comparing
	reconstructed := strings.Join(tokens, "")
	if reconstructed != input {
		return nil, fmt.Errorf("Invalid character(s) in input")
	}

	return tokens, nil
}
