package leetcode

import "strings"

func longestPalindrome(s string) int {
	var (
		read = ""
		maxL = 0
		only = 0
	)

	f := strings.Count(s, string(s[0]))
	if f == len(s) {
		return f
	}

	for _, c := range s {
		if !strings.ContainsRune(read, c) {
			cc := strings.Count(s, string(c))

			if cc == 1 {
				only = 1
			}

			if cc%2 == 1 {
				cc--
				only = 1
			}

			maxL = maxL + cc
			read = read + string(c)
		}
	}
	return maxL + only
}

// esse é mais rapido porem usou mais memoria
func longestPalindromeBestRuntime(s string) int {

	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	foundOdd := false
	length := 0

	for _, char := range freq {
		if char%2 == 0 {
			length += char
		} else {
			length += char - 1
			foundOdd = true
		}
	}

	if foundOdd {
		length++
	}
	return length
}

func longestPalindromeBestRuntimeNewMy(s string) int {
	var (
		foundOdd = false
		maxL     = 0
		read     = ""
	)

	for _, c := range s {
		if !strings.ContainsRune(read, c) {
			length := strings.Count(s, string(c))
			if length%2 == 0 {
				maxL += length
			} else {
				maxL += length - 1
				foundOdd = true
			}
			read = read + string(c)
		}
	}

	if foundOdd {
		maxL++
	}
	return maxL
}
