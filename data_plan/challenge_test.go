/*
CHALLENGE EXPLORATION

plan = 10
months = 5
input = [5, 4, 3, 2, 1]

Month 1.
start = 10
finish = 10 -5 => 5

Month 2.
start = 5 (remaining from previous month) + 10 => 15
finish = 15 - 4 => 11

Month 3.
start = 11 (remaining from previous month) + 10 => 21
finish = 21 - 3 => 18

Month 4.
start = 18 (remaining from previous month) + 10 => 28
finish = 28 - 2 => 26

Month 5.
start = 26 (remaining from previous month) + 10 => 36
finish = 36 - 1 => 35

Output => 35
*/
package main

import (
	"testing"
)

func TestGetAvailableData(t *testing.T) {
	tests := []struct {
		name      string
		usage     []int
		expected  int
		months    int
		plan      int
		expectErr bool
	}{
		{
			name:      "One month no usage",
			plan:      10,
			months:    1,
			usage:     []int{0},
			expected:  20,
			expectErr: false,
		},
		{
			name:      "One month with usage",
			plan:      10,
			months:    1,
			usage:     []int{5},
			expected:  15,
			expectErr: false,
		},
		{
			name:      "Two months no usage",
			plan:      10,
			months:    2,
			usage:     []int{0, 0},
			expected:  30,
			expectErr: false,
		},
		{
			name:      "Two months with usage",
			plan:      10,
			months:    2,
			usage:     []int{5, 4},
			expected:  21,
			expectErr: false,
		},
		{
			name:      "Months and lenght of usage doesn't match",
			plan:      10,
			months:    2,
			usage:     []int{5},
			expected:  21,
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetAvailableData(tc.plan, tc.months, tc.usage)

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
				t.Errorf("Test %q failed: expected %d, got %d", tc.name, tc.expected, got)
			}
		})
	}
}
