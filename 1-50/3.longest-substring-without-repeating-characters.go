package fun50

import "fmt"

func lengthOfLongestSubstring(s string) (ans int) {
	window := [128]bool{}
	left := 0
	for right, c := range s {
		for window[c] {
			window[s[left]] = false
			left++
		}
		window[c] = true
		ans = max(ans, right-left+1)
	}
	return
}

func TestLengthOfLongestSubstring() {
	p := "abcabcbb"
	if res := lengthOfLongestSubstring(p); res != 3 {
		fmt.Printf("TestLengthOfLongestSubstring(%s), returned %d, wanted %d", p, res, 3)
	}
}
