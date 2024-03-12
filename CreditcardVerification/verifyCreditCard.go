package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// to validate the credit cards
func validateCreditCard(number string) string {
	// Define the regex pattern for credit card validation
	pattern := `^[4-6]\d{3}(-?\d{4}){3}$`

	// Compile the regex pattern
	regex := regexp.MustCompile(pattern)

	// Check if the number matches the pattern
	if regex.MatchString(number) {
		return "Valid"
	} else {
		return "Invalid"
	}
}

func main() {
	// Read the number of credit card numbers from input
	var n int
	fmt.Scanln(&n)

	// Read the credit card numbers from input
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n && scanner.Scan(); i++ {
		cardNumber := strings.TrimSpace(scanner.Text())
		fmt.Println(validateCreditCard(cardNumber))
	}
}
