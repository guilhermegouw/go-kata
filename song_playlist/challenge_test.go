package main

import "testing"

func TestGetPlaylist(t *testing.T) {
	tests := []struct {
		name      string
		presses   string
		expected  string
		expectErr bool
	}{
		{
			name:      "One press button 1",
			presses:   "11 41",
			expected:  "B C D E A",
			expectErr: false,
		},
		{
			name:      "One press button 2",
			presses:   "21 41",
			expected:  "E A B C D",
			expectErr: false,
		},
		{
			name:      "One press button 3",
			presses:   "31 41",
			expected:  "B A C D E",
			expectErr: false,
		},
		{
			name:      "One press button 4",
			presses:   "41",
			expected:  "A B C D E",
			expectErr: false,
		},
		{
			name:      "Two presses button 1",
			presses:   "12 41",
			expected:  "C D E A B",
			expectErr: false,
		},
		{
			name:      "Two presses button 2",
			presses:   "22 41",
			expected:  "D E A B C",
			expectErr: false,
		},
		{
			name:      "Two presses button 3",
			presses:   "32 41",
			expected:  "A B C D E",
			expectErr: false,
		},
		{
			name:      "Multiple presses of button 4",
			presses:   "44",
			expected:  "A B C D E",
			expectErr: false,
		},
		{
			name:      "Invalid input pair with one char instead of two",
			presses:   "1 41",
			expected:  "",
			expectErr: true,
		},
		{
			name:      "Invalid input pair with three chars instead of two",
			presses:   "321 41",
			expected:  "",
			expectErr: true,
		},
		{
			name:      "Invalid button value",
			presses:   "x2 41",
			expected:  "",
			expectErr: true,
		},
		{
			name:      "Invalid number of presses",
			presses:   "3x 41",
			expected:  "",
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := GetPlaylist(tc.presses)

			if tc.expectErr {
				if err == nil {
					t.Errorf("Test %q failed: expected an error but got none", tc.name)
				}
			} else if err != nil {
				t.Errorf("Test %q failed: unexpected error: %v", tc.name, err)
			} else if got != tc.expected {
				t.Errorf("Test %q failed: expected %q, got %q", tc.name, tc.expected, got)
			}
		})
	}
}
