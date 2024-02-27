package main

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var testCases = []struct {
		input    string
		expected bool
	}{
		{"", false},
		{" ", false},
		{"hello", true},
		{" hello", true},
		{"hello ", true},
	}

	for _, tc := range testCases {
		if ok := IsEmpty(tc.input); ok != tc.expected {
			t.Errorf("IsEmpty (%q) expected %v, but got %v.", tc.input, tc.expected, ok)
		}
	}
}

func TestCheckLength(t *testing.T) {
	var testCases = []struct {
		input  string
		result error
	}{
		{input: "xjda-43a", result: nil},
		{input: "password", result: nil},
		{input: "pass", result: TooShortErr},
		{input: "", result: TooShortErr},
		{input: "thisIsAVeryLongInputThatExceedsTheMaxLengthAllowed", result: TooLongErr},
		{input: "xjda-43as-eeja-n8vv12345678901234", result: TooLongErr},
	}

	for _, tc := range testCases {
		if err := CheckLength(tc.input); err != nil && err != tc.result {
			t.Errorf("CheckLength (%q) expected %v, but got %v.", tc.input, tc.result, err)
		} else if err == nil && tc.result != nil {
			t.Errorf("CheckLength (%q) expected error %v, but got nil.", tc.input, tc.result)
		}
	}
}

func TestIsMatching(t *testing.T) {
	var testCases = []struct {
		a, b   string
		result bool
	}{
		{a: "is", b: "is", result: true},
		{a: " ", b: " ", result: true},
		{a: " hello", b: "hello ", result: false},
		{a: "hello", b: "HelLo ", result: false},
		{a: "129nvbn", b: "1193nljs ", result: false},
		{a: "123", b: "123", result: true},
		{a: "abc", b: "ABC", result: false},
	}

	for _, tc := range testCases {
		if ok := IsMatching(tc.a, tc.b); ok != tc.result {
			t.Errorf("IsMatching (%q, %q) expected %v, but got %v.", tc.a, tc.b, tc.result, ok)
		}
	}
}

func TestCheckEmail(t *testing.T) {
	var testCases = []struct {
		input string 
		expected bool
	}{
		{input: "ezlos com", expected: false},
		{input: "bad-example", expected: false},
		{input: "ezlos@example.com", expected: true},
		{input: "good@example.com", expected: true},
	}

	for _, tc := range testCases {
		if ok := CheckEmail(tc.input); ok != tc.expected {
			t.Errorf("CheckEmail (%q) expected %v, but got %v.", tc.input, tc.expected, ok)
		}
	}
}
