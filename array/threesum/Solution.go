package threesum

import "sort"

func threeSum(nums []int) (value [][]int) {
	sort.Ints(nums)
	if len(nums) == 3 && (nums[0]+nums[1]+nums[2] == 0) {
		value = append(value, nums)
		return
	}

	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || nums[i] > nums[i+1] {
			start := i + 1
			end := len(nums) - 1
			for {
				sum := nums[i] + nums[start] + nums[end]
				if sum == 0 {
					value = append(value, []int{nums[i], nums[start], nums[end]})
				}
				if sum > 3 {
					temp := nums[end]
					for {
						end--
						if end == start || temp != nums[end] {
							break
						}
					}
				} else {
					temp := nums[start]
					for {
						start++
						if start == end || temp != nums[start] {
							break
						}
					}
				}
				if start == end {
					break
				}
			}
		}
	}

	return
}
