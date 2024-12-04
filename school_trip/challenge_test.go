package main

import (
	"testing"
)

func TestNeedsMoreMoney(t *testing.T) {
	tests := []struct {
		name        string
		cost        int
		proportions [4]float64
		students    int
		expected    string
		expectErr   bool
	}{
		{
			name:        "4 students one of each year.",
			cost:        100,
			proportions: [4]float64{0.25, 0.25, 0.25, 0.25},
			students:    4,
			expected:    "YES",
			expectErr:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := NeedsMoreMoney(tc.cost, tc.proportions, tc.students)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Test %q failed: expected and error but got none", tc.name)
				}
				return
			}

			if err != nil {
				t.Errorf("Test %q failed: unexpected error: %v", tc.name, err)
				return
			}

			if got != tc.expected {
				t.Errorf("Test %q failed: expected %s, got %s", tc.name, tc.expected, got)
			}
		})
	}
}

/*
EXPLORING THE CHALLENGE:

Case 1:

first year: 4 * 0.25 = 1 => 12
second year: 4 * 0.25 = 1 => 10
third year: 4 * 0.25 = 1 => 7
fourth year: 4 * 0.25 = 1 => 5

Total raised: 12 + 10 + 7 + 5 => 34
needs more money? cost=100 - total=34 => 66
result: YES
*/
