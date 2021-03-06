package main

import "fmt"


//var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubStr(s string) int {
	// stores last occurred pos + 1
	// 0 means not seen.

	//for i := range lastOccurred {
	//	lastOccurred[i] = 0
	//}
	//
	//start := 0
	//maxLength := 0
	//
	//for i, ch := range []rune(s) {
	//	if lastI := lastOccurred[ch]; lastI >= start {
	//		start = lastI
	//	}
	//	if i - start + 1 > maxLength {
	//		maxLength = i - start + 1
	//	}
	//	lastOccurred[ch] = i + 1
	//}
	//
	//return maxLength

	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcabc"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("asdfghjklzxc"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("这里是慕课网"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("一二三二一"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
