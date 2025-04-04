package dsa

func MajorityElement(nums []int) int {
	count := 1
	ele := nums[0]

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			ele = nums[i]
		}
		if ele == nums[i] {
			count++
		} else {
			count--
		}
	}
	return ele
}

// [2,2,1,1,1,2,2]
// [3,2,3]
