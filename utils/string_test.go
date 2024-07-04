package utils

import (
    "testing"
)

type stringTestCase struct {
    name     string
    input    string
    expected int
}

func TestGetLength(t *testing.T) {
    testCases := []stringTestCase{
        {"Empty string", "", 0},
        {"Single character", "a", 1},
        {"Multiple characters", "hello world", 11},
        {"Whitespace only", "   ", 3},
        {"Special characters", "!@#$%^&*()", 10},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            actual := GetLength(tc.input)
            if actual != tc.expected {
                t.Errorf("Expected %d, got %d", tc.expected, actual)
            }
        })
    }
}

