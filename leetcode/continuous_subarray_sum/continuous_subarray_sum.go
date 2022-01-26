// https://leetcode.com/problems/continuous-subarray-sum/
package continuous_subarray_sum

func checkSubarraySum(nums []int, k int) bool {
	for cnt := 2; cnt <= len(nums); cnt++ {
		for i := 0; i <= len(nums)-cnt; i++ {
			sum := getSum(nums, i, cnt)
			if (k == 0 && sum == 0) || (k != 0 && sum%k == 0) {
				return true
			}
		}
	}
	return false
}

var sums map[int]int = make(map[int]int)

func getSum(nums []int, i, cnt int) int {
	sum := 0
	if cnt == 2 {
		sum = nums[i] + nums[i+1]
		sums[i] = sum
	} else {
		v, _ := sums[i]
		sum = nums[i+cnt-1] + v
		sums[i] = sum
	}
	return sum
}
