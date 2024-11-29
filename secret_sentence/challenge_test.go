package main

import "testing"

func TestLukaEncoder(t *testing.T) {
	tests := []struct {
		name     string
		decoded  string
		expected string
	}{
		{
			name:     "One consonant",
			decoded:  "b",
			expected: "b",
		},
		{
			name:     "Two consonants",
			decoded:  "bb",
			expected: "bb",
		},
		{
			name:     "One vowel",
			decoded:  "a",
			expected: "apa",
		},
		{
			name:     "Two different vowels",
			decoded:  "ae",
			expected: "apaepe",
		},
		{
			name:     "Two equal vowels",
			decoded:  "aa",
			expected: "apaapa",
		},
		{
			name:     "Two different vowels with space",
			decoded:  "a e",
			expected: "apa epe",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LukaEncoder(tc.decoded)

			if tc.expected != got {
				t.Errorf("Test %q failed: expected %s, got %s", tc.name, tc.expected, got)
			}
		})
	}
}

func TestLukaDecoder(t *testing.T) {
	tests := []struct {
		name     string
		encoded  string
		expected string
	}{
		{
			name:     "One consonant",
			encoded:  "b",
			expected: "b",
		},
		{
			name:     "Two consonants",
			encoded:  "bb",
			expected: "bb",
		},
		{
			name:     "One vowel",
			encoded:  "apa",
			expected: "a",
		},
		{
			name:     "Two different vowels",
			encoded:  "apaepe",
			expected: "ae",
		},
		{
			name:     "Two equal vowels",
			encoded:  "apaapa",
			expected: "aa",
		},
		{
			name:     "Two different vowels with space",
			encoded:  "apa epe",
			expected: "a e",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := LukaDecoder(tc.encoded)

			if tc.expected != got {
				t.Errorf("Test %q failed: expected %s, got %s", tc.name, tc.expected, got)
			}
		})
	}
}
