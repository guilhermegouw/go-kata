package main

import (
	"testing"
)

func TestCountPlays(t *testing.T) {
	tests := []struct {
		name          string
		quarters      int
		firstMachine  int
		secondMachine int
		thirdMachine  int
		expected      int
	}{
		{
			name:          "Martha wins on first machine at first play",
			quarters:      1,
			firstMachine:  34,
			secondMachine: 0,
			thirdMachine:  0,
			expected:      40,
		},
		{
			name:          "Martha wins on second machine at second play",
			quarters:      2,
			firstMachine:  0,
			secondMachine: 99,
			thirdMachine:  0,
			expected:      80,
		},
		{
			name:          "Martha wins on third machine at third play",
			quarters:      3,
			firstMachine:  0,
			secondMachine: 0,
			thirdMachine:  9,
			expected:      12,
		},
		{
			name:          "Single quarter no wins",
			quarters:      1,
			firstMachine:  0,
			secondMachine: 0,
			thirdMachine:  0,
			expected:      1,
		},
		{
			name:          "Maximum quarters (1000) starting fresh",
			quarters:      1000,
			firstMachine:  0,
			secondMachine: 0,
			thirdMachine:  0,
			expected:      4375,
		},
		{
			name:          "All machines about to pay",
			quarters:      5,
			firstMachine:  34,
			secondMachine: 99,
			thirdMachine:  9,
			expected:      179,
		},
		{
			name:          "No initial quarters",
			quarters:      0,
			firstMachine:  0,
			secondMachine: 0,
			thirdMachine:  0,
			expected:      0,
		},
		{
			name:          "Third machine frequent wins scenario",
			quarters:      10,
			firstMachine:  0,
			secondMachine: 0,
			thirdMachine:  8,
			expected:      19,
		},
		{
			name:          "Almost winning combinations",
			quarters:      5,
			firstMachine:  33,
			secondMachine: 98,
			thirdMachine:  8,
			expected:      179,
		},
		{
			name:          "Multiple consecutive wins possible",
			quarters:      3,
			firstMachine:  34,
			secondMachine: 99,
			thirdMachine:  9,
			expected:      177,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CountPlays(tc.quarters, tc.firstMachine, tc.secondMachine, tc.thirdMachine)

			if got != tc.expected {
				t.Errorf("Test %q failed: expected %d, got %d", tc.name, tc.expected, got)
			}
		})
	}
}
