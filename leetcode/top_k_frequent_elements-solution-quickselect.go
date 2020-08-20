// https://leetcode.com/problems/top-k-frequent-elements

// Based on algorithm: https://en.wikipedia.org/wiki/Quickselect
package main

func topKFrequent(nums []int, k int) []int {
	// O(N)
	fs := freqs(nums)

	// O(N)
	frV := make([]int, len(fs))
	elV := make([]int, len(fs))
	idx := 0
	for eV, fV := range fs {
		elV[idx] = eV
		frV[idx] = fV
		idx++
	}

	// quickselect - O(N), the worst case will be O(N^2) and we need Hoare partitioning to work with duplicates properly
	quickselect(elV, frV, 0, len(fs)-1, len(fs)-k)
	return elV[len(elV)-k:]
}

// freqs - returns map[elementValue]frequency(count of elements)
func freqs(nums []int) map[int]int {
	res := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		f, _ := res[nums[i]]
		res[nums[i]] = f + 1
	}
	return res
}

func quickselect(elV, frV []int, left, right, k int) {
	for {
		if left == right {
			return
		}
		pivotIndex := partitionLomuto(elV, frV, left, right)
		if pivotIndex == k {
			return
		}
		if k < pivotIndex {
			quickselect(elV, frV, left, pivotIndex-1, k)
			return
		} else {
			quickselect(elV, frV, pivotIndex+1, right, k)
			return
		}
	}
}

func quicksort(elV, frV []int, left, right int) {
	if left < right {
		pivotIndex := partitionLomuto(elV, frV, left, right)
		quicksort(elV, frV, left, pivotIndex-1)
		quicksort(elV, frV, pivotIndex+1, right)
	}
}

// Hoare partition
func partitionHoare(elV, frV []int, left, right int) int {
	pivotIndex := (right + left) / 2
	pivot := frV[pivotIndex]
	for {
		for frV[left] < pivot {
			left++
		}
		for frV[right] > pivot {
			right--
		}

		if left >= right {
			return right
		}

		// swap
		swap(frV, left, right)
		swap(elV, left, right)
		left++
		right--
	}
}

// Lomuto partition
func partitionLomuto(elV, frV []int, left, right int) int {
	pivot := frV[right]
	i := left
	for j := left; j <= right; j++ {
		if frV[j] < pivot {
			swap(frV, i, j)
			swap(elV, i, j)
			i++
		}
	}
	swap(frV, i, right)
	swap(elV, i, right)
	return i
}

func swap(n []int, i, j int) {
	t := n[i]
	n[i] = n[j]
	n[j] = t
}
