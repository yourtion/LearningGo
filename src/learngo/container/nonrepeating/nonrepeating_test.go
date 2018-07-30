package main

import "testing"

func TestSubstr(t *testing.T) {

	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"abaabcbb", 3},
		{"bbbbb", 1},
		{"pwwkwe", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"测试测试", 2},
		{"一二三一二", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; expected %d", tt.s, actual, tt.ans)
		}
	}
}
