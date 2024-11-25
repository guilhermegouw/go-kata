package challenge

import (
	"testing"
)

func TestOccupiedOnBothDays(t *testing.T) {
	tests := []struct {
		name      string
		numSpaces int
		yesterday string
		today     string
		expected  int
		expectErr bool
	}{
		{
			name:      "No cars on both days",
			numSpaces: 1,
			yesterday: ".",
			today:     ".",
			expected:  0,
			expectErr: false,
		},
		{
			name:      "1 car on one day and, none the other",
			numSpaces: 1,
			yesterday: "C",
			today:     ".",
			expected:  0,
			expectErr: false,
		},
		{
			name:      "1 car on both days",
			numSpaces: 1,
			yesterday: "C",
			today:     "C",
			expected:  1,
			expectErr: false,
		},
		{
			name:      "1 car on both days 0 on the other",
			numSpaces: 2,
			yesterday: "C.",
			today:     "C.",
			expected:  1,
		},
		{
			name:      "2 spaces in 2 days",
			numSpaces: 2,
			yesterday: "CC",
			today:     "CC",
			expected:  2,
		},
		{
			name:      "Mismatched yesterday and today input lengths",
			numSpaces: 1,
			yesterday: "C.",
			today:     "C",
			expected:  0,
			expectErr: true,
		},
		{
			name:      "Mismatched numSpaces with yesterday and today input lengths",
			numSpaces: 1,
			yesterday: "C.",
			today:     "CC",
			expected:  0,
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := OccupiedOnBothDays(tc.numSpaces, tc.yesterday, tc.today)

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
