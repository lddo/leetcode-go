package main

import (
	"fmt"
	"leetcode/algorithm"
)

func main() {
	// 1. 两数之和
	//nums := []int{2, 7, 11, 15}
	//target := 9
	//fmt.Println(algorithm.TwoSum(nums, target))
	//fmt.Println(algorithm.TwoSumForHash(nums, target))

	// 2. 两数相加
	//ln1 := &algorithm.ListNode{
	//	Val: 9,
	//	Next: &algorithm.ListNode{
	//		Val: 9,
	//		Next: &algorithm.ListNode{
	//			Val:  9,
	//			Next: nil,
	//		},
	//	},
	//}
	//ln2 := &algorithm.ListNode{
	//	Val: 9,
	//	Next: &algorithm.ListNode{
	//		Val: 9,
	//	},
	//}
	//fmt.Println(algorithm.PrintNumbers(algorithm.AddTwoNumbers(ln1, ln2)))

	// 3. 无重复字符的最长子串
	//fmt.Println(algorithm.LengthOfLongestSubstring("abcabcbb"))

	// 5. 最长回文子串
	fmt.Println(algorithm.LongestPalindrome("a"))
	fmt.Println(algorithm.LongestPalindrome("ac"))
	fmt.Println(algorithm.LongestPalindrome("babad"))
	fmt.Println(algorithm.LongestPalindrome("abca"))
	fmt.Println(algorithm.LongestPalindrome("aaaaa"))
	fmt.Println(algorithm.LongestPalindrome("ababcabcda"))
	fmt.Println(algorithm.LongestPalindrome("ccc"))
	fmt.Println(algorithm.LongestPalindrome("aacabdkacaa"))

	// 计算执行时间
	//start := time.Now().UnixMicro()
	//fmt.Println(time.Now().UnixMicro() - start)

	str := ""
	fmt.Println(str[:0])
}
