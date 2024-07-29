package main

// You and Fredrick are good friends. Yesterday, Fredrick received  credit cards from ABCD Bank. He wants to verify whether his credit card numbers are valid or not. You happen to be great at regex so he is asking for your help!

// A valid credit card from ABCD Bank has the following characteristics:

// ► It must start with a ,  or .
// ► It must contain exactly  digits.
// ► It must only consist of digits (-).
// ► It may have digits in groups of , separated by one hyphen "-".
// ► It must NOT use any other separator like ' ' , '_', etc.
// ► It must NOT have  or more consecutive repeated digits.

// Examples:

// Valid Credit Card Numbers

// 4253625879615786
// 4424424424442444
// 5122-2368-7954-3214
// Invalid Credit Card Numbers

// 42536258796157867       #17 digits in card number → Invalid
// 4424444424442444        #Consecutive digits are repeating 4 or more times → Invalid
// 5122-2368-7954 - 3214   #Separators other than '-' are used → Invalid
// 44244x4424442444        #Contains non digit characters → Invalid
// 0525362587961578        #Doesn't start with 4, 5 or 6 → Invalid
// Input Format

// The first line of input contains an integer .
// The next  lines contain credit card numbers.

// Constraints

// Output Format

// Print 'Valid' if the credit card number is valid. Otherwise, print 'Invalid'. Do not print the quotes.

// Sample Input

// 6
// 4123456789123456
// 5123-4567-8912-3456
// 61234-567-8912-3456
// 4123356789123456
// 5133-3367-8912-3456
// 5123 - 3567 - 8912 - 3456
// Sample Output

// Valid
// Valid
// Invalid
// Valid
// Invalid
// Invalid
// Explanation

// 4123456789123456 : Valid
// 5123-4567-8912-3456 : Valid
// 61234--8912-3456 : Invalid, because the card number is not divided into equal groups of .
// 4123356789123456 : Valid
// 51-67-8912-3456 : Invalid, consecutive digits  is repeating  times.
// 5123456789123456 : Invalid, because space '  ' and - are used as separators.

import (
	"fmt"
	"regexp"
	"strings"
)

func isValidCreditCard(cardNumber string) bool {
	// Regex pattern for valid credit card
	// Explanation:
	// ^                     : Start of string
	// (	                 : Start of group
	// [4-6]                 : Start with digit 4, 5, or 6
	// (                     : Start of group
	//   [0-9]{16}           : Exactly 16 digits
	//   |                   : OR
	// (					 : Start of group
	// [4-6]                 : Start with digit 4, 5, or 6
	// [0-9]{3}-)		     : 3 digits followed by a hyphen
	// ){1}                  : End of  one group
	// (					 : Start of group
	// [0-9]{4}-){2}         : End of two groups of 4 digits followed by a hyphen
	//   [0-9]{4}            : Final group of 4 digits
	// )                     : End of group
	// $                     : End of string
	//pattern := `^([4-6]|[4-6][0-9]{15}|(([4-6]([0-9]{3}-){1}([0-9]{4}-){2}[0-9]{4}))$`
	pattern := `^([4-6][0-9]{15}|([4-6]([0-9]{3}-){1}([0-9]{4}-){2}[0-9]{4}))$`
	re := regexp.MustCompile(pattern)

	// First, check if the card number matches the regex pattern
	if !re.MatchString(cardNumber) {
		return false
	}

	// Remove hyphens for checking consecutive digits
	cardNumber = strings.ReplaceAll(cardNumber, "-", "")

	// Check for four consecutive identical digits
	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i] == cardNumber[i+2] && cardNumber[i] == cardNumber[i+3] {
			return false
		}
	}

	return true
}

func main() {
	cards := []string{
		"4253625879615786",          // Valid
		"4424424424442444",          // Valid
		"5122-2368-7954-3214",       // Valid (group length incorrect)
		"42536258796157867",         // 17 digits in card number → Invalid
		"4424444424442444",          // Consecutive digits are repeating 4 or more times → Invalid
		"5122-2368-7954 - 3214",     // Separators other than '-' are used → Invalid
		"44244x4424442444",          // Contains non digit characters → Invalid
		"0525362587961578",          // Doesn't start with 4, 5 or 6 → Invalid
		"4123456789123456",          // Valid
		"5123-4567-8912-3456",       // Valid
		"61234-567-8912-3456",       // Invalid, because the card number is not divided into equal groups of 4.
		"4123356789123456",          // Valid
		"5133-3367-8912-3456",       // Invalid, consecutive digits 3333 is repeating 4 times.
		"5123 - 3567 - 8912 - 3456", // Invalid, because space ' ' and - are used as separators.
	}

	for _, card := range cards {
		if isValidCreditCard(card) {
			fmt.Printf("%s: Valid\n", card)
		} else {
			fmt.Printf("%s: Invalid\n", card)
		}
	}
}

// go run main.go
