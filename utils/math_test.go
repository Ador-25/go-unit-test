package utils

import (
	"testing"
)

type mathTestCase struct {
	name     string
	a        int
	b        int
	expected int
}

func TestAdd(t *testing.T) {
	testCases := []mathTestCase{
		{"Case 1", 2, 3, 5},
		{"Case 2", -1, 1, 0},
		{"Case 3", 10, 20, 30},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Fatalf("%s: %d + %d != %d. expected = %d", tc.name, tc.a, tc.b, result, tc.expected)
		} else {
			t.Logf("%s: Passed", tc.name)
		}
	}
}

func TestSub(t *testing.T) {
	testCases := []mathTestCase{
		{"Case 1", 2, 3, -1},
		{"Case 2", -1, 1, -2},
		{"Case 3", 10, 20, -10},
	}

	for _, tc := range testCases {
		result := Sub(tc.a, tc.b)
		if result != tc.expected {
			t.Fatalf("%s: %d - %d != %d. expected = %d", tc.name, tc.a, tc.b, result, tc.expected)
		} else {
			t.Logf("%s: Passed", tc.name)
		}
	}
}
