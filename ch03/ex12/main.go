package main

import "fmt"

func string2byteCount(s string) map[byte]int {
	cnt := map[byte]int{}
	for _, b := range []byte(s) {
		cnt[b]++
	}
	return cnt
}

func isAnagram(x, y string) bool {
	if x == y {
		return true
	}

	if len(x) != len(y) {
		return false
	}

	xByteCnt := string2byteCount(x)
	yByteCnt := string2byteCount(y)
	if len(xByteCnt) != len(yByteCnt) {
		return false
	}

	for xbKey, xbCnt := range xByteCnt {
		ybCnt, ok := yByteCnt[xbKey]
		if !ok {
			return false
		}
		if xbCnt != ybCnt {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("abcd", "adcb"))
	fmt.Println(isAnagram("abcd", "aaa"))
	fmt.Println(isAnagram("abcd", "aaab"))
	fmt.Println(isAnagram("abcd", "abcd"))
	fmt.Println(isAnagram("あん", "んあ"))
}
