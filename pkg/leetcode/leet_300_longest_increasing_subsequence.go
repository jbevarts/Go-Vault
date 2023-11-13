package leetcode

// lengthOfLIS returns length of longest, increasing subsequence of nums.
func lengthOfLIS(nums []int) int {

	longestSoFar := []int{}
	for range nums {
		longestSoFar = append(longestSoFar, 0)
	}
	maxSoFar := 0

	for i, v := range nums {
		bestYet := 0
		for k, w := range nums[:i] {
			if w < v && longestSoFar[k] > bestYet {
				bestYet = longestSoFar[k]
			}
		}
		longestSoFar[i] = 1 + bestYet
		if longestSoFar[i] > maxSoFar {
			maxSoFar = longestSoFar[i]
		}
	}
	return maxSoFar
}
