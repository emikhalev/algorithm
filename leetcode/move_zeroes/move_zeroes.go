// https://leetcode.com/problems/move-zeroes/
package move_zeroes

func moveZeroes(nums []int) {
	zIdx, nzIdx := 0, 0

ext:
	for nzIdx < len(nums) && zIdx < len(nums) {

		for nums[zIdx] != 0 {
			zIdx++
			if zIdx >= len(nums) {
				break ext
			}
		}
		for nums[nzIdx] == 0 {
			nzIdx++
			if nzIdx >= len(nums) {
				break ext
			}
		}

		if zIdx < nzIdx {
			nums[zIdx] = nums[nzIdx]
			nums[nzIdx] = 0

			zIdx++
			nzIdx = zIdx + 1
		} else {
			nzIdx++
			zIdx = nzIdx
		}

	}
}
