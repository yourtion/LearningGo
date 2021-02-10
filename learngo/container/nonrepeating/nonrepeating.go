package main

import "fmt"

var lastOccurred = make([]int, 0xfffff)

func lengthOfNonRepeatingSubStr(s string) int {

	// stores last occurred pos + 1.
	// 0 means not seen.
	start := 0
	maxLength := 0

	for i := range lastOccurred {
		lastOccurred[i] = 0
	}

	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI > start {
			start = lastOccurred[ch]
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i + 1
	}

	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abaabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkwe"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdef"))

	fmt.Println(
		lengthOfNonRepeatingSubStr("测试测试"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三一二"))
}
