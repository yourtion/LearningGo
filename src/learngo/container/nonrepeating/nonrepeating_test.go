package main

import "testing"

// go test -coverprofile=c.out
// go tool cover -html=c.out

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
		{"黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; expected %d", tt.s, actual, tt.ans)
		}
	}
}

//  go test -bench .

func BenchmarkSubstr(b *testing.B) {

	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("lengthOfNonRepeatingSubStr(%s); got %d; expected %d", s, actual, ans)
		}
	}
}
