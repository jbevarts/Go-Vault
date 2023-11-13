package leetcode

import "fmt"

// Given an array of integers `nums` and an integer
// `k` return the total number of subarrays whose sum equals to k.
// A subarray is a contiguous non-empty sequence of elements within an array.
// constraints:
// 1 <= nums.length <= 2 * 104
// -1000 <= nums[i] <= 1000
// -107 <= k <= 107

// Initial Idea:
// Total time is: 2*len(nums) = O(n)
// Start at element 0 with two pointers, A and B.
// Scoot the B pointer along until it reaches a subarray sum greater than or equal to k.
// Scoot A towards B until the sum is again greater or equal.

// must consider negative numbers, challenging when doing sliding window with two pointers beecause recursive in nature.
// cannot normalize values because array can be too large.

// prefix sum array
// Sum of arrays ending at i, stored in hashmap.
// Then, iterate through map counting all of the exact matches.

// create map
// for each element, store a mapping of int->[]int containing possible sums leading.

// ideas:
// - there can be multiple solutions ending at the same element. [2, -2, 2, -2] ~ all 4 is same as last 2.
// - there can be multiple solutions starting at same element. " "

// o(n log n)
// n = full traverse
// log n = binary search for match

func prefixSumsWithExactMatchCount(nums []int, k int) []int {
	var totalSoFar int
	numsSum := make([]int, len(nums))
	for i, v := range nums {
		numsSum[i] = totalSoFar + v
		totalSoFar = numsSum[i]
	}
	return numsSum
}

func subarraySum(nums []int, k int) int {
	prefixSums := prefixSumsWithExactMatchCount(nums, k)
	// fmt.Println(prefixSums, matches)
	counter := 0

	// we want to find the number of j's, that is greater than i, where nums[j] - nums[i] = k
	// need to avoid redundant solutions.

	for i, n := range prefixSums {
		// nums[i] is the prefix sum, an exact match should increment.
		if nums[i] == k {
			counter += 1
		}
		for _, m := range prefixSums {
			// fmt.Println(m, n)
			if m-n == k {
				fmt.Println("here")
				counter += 1
			}
		}
	}
	return counter
}
