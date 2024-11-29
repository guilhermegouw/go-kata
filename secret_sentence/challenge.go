/*
The Challenge

Luka is writing a secret sentence in class.
He doesn’t want the teacher to be able to read it, so instead of writing down
the original sentence, he writes down an encoded version.
After each vowel in the sentence (a, e, i, o, or u), he adds the letter p and
that vowel again.

For example, rather than write down the sentence
i like you

he would write:
ipi lipikepe yopoupu.

The teacher acquires Luka’s encoded sentence.
Recover Luka’s original sentence for the teacher.

# Input

The input is one line of text, Luka’s encoded sentence.
It consists of lowercase letters and spaces.
There is exactly one space between each pair of words.

# Output

Output Luka’s original sentence.
*/
package main

import (
	"strings"
)

func isVowel(letter rune) bool {
	return strings.ContainsRune("aeiou", letter)
}

func LukaEncoder(decoded string) string {
	var builder strings.Builder
	for _, letter := range decoded {
		if isVowel(letter) {
			builder.WriteRune(letter)
			builder.WriteRune('p')
			builder.WriteRune(letter)
		} else {
			builder.WriteRune(letter)
		}
	}
	return builder.String()
}

func LukaDecoder(encoded string) string {
	var builder strings.Builder
	for i := 0; i < len(encoded); i++ {
		char := rune(encoded[i])
		if isVowel(char) {
			builder.WriteRune(char)
			i += 2
		} else {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}
