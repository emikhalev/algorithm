// https://leetcode.com/problems/top-k-frequent-elements
package main

func topKFrequent(nums []int, k int) []int {
	freqs := make(map[int]int)         // frequencies of all elements
	topFreqs := make(map[int]struct{}) // top K frequencies

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		// update freqs
		cnt, _ := freqs[num]
		cnt++
		freqs[num] = cnt

		// if not if topFreqs
		_, ok := topFreqs[num]
		if !ok {
			if len(topFreqs) < k { // If topFreq have empty space - just add
				topFreqs[num] = struct{}{}
			} else { // if topFreq is already filled - compare with min element and replace if needed
				minNum, minCnt := getMinCnt(freqs, topFreqs)
				if cnt > minCnt {
					delete(topFreqs, minNum)
					topFreqs[num] = struct{}{}
				}
			}
		}
	}
	ans := make([]int, k)
	idx := 0
	for num, _ := range topFreqs {
		ans[idx] = num
		idx++
	}
	return ans
}

func getMinCnt(freqs map[int]int, minFreqs map[int]struct{}) (minNum int, minCnt int) {
	minCnt = 1<<32 - 1
	minNum = -1
	for num, _ := range minFreqs {
		cnt, _ := freqs[num]
		if minCnt > cnt {
			minNum = num
			minCnt = cnt
		}
	}
	return minNum, minCnt
}
