// https://leetcode.com/problems/coin-change/
package coin_change

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	return coinsCnt(coins, amount, make(map[int]int))
}

func coinsCnt(coins []int, sum int, count map[int]int) int {
	if sum < 0 {
		return -1
	}
	if sum == 0 {
		return 0
	}
	if v, ok := count[sum-1]; ok {
		return v
	}

	minCnt := 1<<32 - 1
	for i := 0; i < len(coins); i++ {
		cv := coins[i]
		cnt := coinsCnt(coins, sum-cv, count)
		if cnt >= 0 && cnt < minCnt {
			minCnt = 1 + cnt
		}
	}

	if minCnt == 1<<32-1 {
		minCnt = -1
	}
	count[sum-1] = minCnt
	return minCnt
}
