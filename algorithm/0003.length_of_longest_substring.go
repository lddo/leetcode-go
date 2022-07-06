package algorithm

func LengthOfLongestSubstring(s string) int {
	sLen := len(s)
	bMap := map[byte]int{}
	left, maxLen := 0, 0
	for i := 0; i < sLen; i++ {
		b := s[i]
		if j, ok := bMap[b]; ok {
			for ; left < j+1; left++ {
				delete(bMap, s[left])
			}
		}
		bMap[b] = i
		maxLen = max(maxLen, i+1-left)
	}
	return maxLen
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
