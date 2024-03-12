package leetcode

type T struct{}

func ContainsNearbyDuplicate(nums []int, k int) bool {
	hs := make(map[int]T)
	for i := 0; i < len(nums); i++ {
		_, isOk := hs[nums[i]]
		if isOk {
			return true
		}
		hs[nums[i]] = T{}
		if i >= k {
			delete(hs, nums[i-k])
		}
	}
	return false
}
