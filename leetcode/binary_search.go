package leetcode

func Search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (high + low) / 2
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		}
	}
	return -1
}
