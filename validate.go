package validator

import (
	"errors"
	"net/mail"
	"strings"
)

var (
	// length
	min = 8
	max = 32

	// errors
	TooShortErr = errors.New("input provided is too short")
	TooLongErr  = errors.New("input provided is too long")
)

// IsEmpty will return false if no characters were provided in the input.
func IsEmpty(input string) bool {
	input = strings.TrimSpace(input)
	return input != ""
}

// CheckLength checks for the length of the input and returns an error if
// the length of the input is less than 8 or larger than 32.
func CheckLength(input string) error {
	if len(input) < min {
		return TooShortErr
	}

	if len(input) > max {
		return TooLongErr
	}

	return nil
}

// IsMatching checks if the both values provided are equal. If both values
// are equal, it returns true.
func IsMatching(a, b string) bool {
	return a == b
}

// CheckEmail checks if the provided email is valid. If it is valid, CheckEmail
// returns true.
func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}
