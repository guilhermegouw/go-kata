/*
Exploring the problem:

Case 3 villages:
|first|second|third|
|  6  |  10  |  15 |

second village starts at: 10 + 6 = 16 => 16/2 = 8 (position)
second village ends at: 10 + 15 = 25 => 25/2 = 12.5 (position)
second village size: 12.5 - 8 = 4.5

Case 4 villages:
|first|second|third|fourth|
|  6  |  10  |  15 |  19  |

second village starts at: 10 + 6 = 16 => 16/2 = 8 (position)
second village ends at: 10 + 15 = 25 => 25/2 = 12.5 (position)
second village size: 12.5 - 8 = 4.5

third village starts at: 10 + 15 = 25 => 25/2 = 12.5 (position)
third village ends at: 15 + 19 = 34 => 34/2 = 17 (position)
third village size: 17 - 12.5 = 4.5

size of the smallest village = 4.5
*/
package main

import (
	"testing"
)

func TestGetSmallestNeighborhood(t *testing.T) {
	tests := []struct {
		name        string
		numVillages int
		posVillages []int
		expected    float32
		expectErr   bool
	}{
		{
			name:        "3 villages {6, 10, 15}",
			numVillages: 3,
			posVillages: []int{6, 10, 15},
			expected:    4.5,
			expectErr:   false,
		},
		{
			name:        "4 villages {6, 10, 15, 18}",
			numVillages: 4,
			posVillages: []int{6, 10, 15, 18},
			expected:    4.0,
			expectErr:   false,
		},
		{
			name:        "Invalid: Too few villages",
			numVillages: 2,
			posVillages: []int{1, 2},
			expected:    0,
			expectErr:   true,
		},
		{
			name:        "Invalid: Count mismatch",
			numVillages: 4,
			posVillages: []int{1, 2, 3},
			expected:    0,
			expectErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetSmallestNeighborhood(tc.numVillages, tc.posVillages)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Test %q failed: expected and error but got none", tc.name)
				}
				return
			}
			if err != nil {
				t.Errorf("Test %q failed: unexpected error: %v", tc.name, err)
			}
			if got != tc.expected {
				t.Errorf("Test %q failed: expected %.1f, got %.1f", tc.name, tc.expected, got)
			}
		})
	}
}
