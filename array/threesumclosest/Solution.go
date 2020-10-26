package main

import (
	"fmt"
	"math"
	"sort"
)

//[-1,3,2,-4,5,1] == 4
//[-4,-1,1,2,3,5]

func main() {
	fmt.Println(threeSumClosest([]int{-1,3,2,-4,5,1}, 4))
}

func threeSumClosest(nums []int, target int) (result int) {
	sort.Ints(nums)
	closestSum := math.MaxInt64

	for i := 0; i < len(nums) && closestSum != 0; i++ {
		start := i + 1
		end := len(nums) - 1
		for {
			if start >= end {
				break
			}
			sum := nums[i] + nums[start] + nums[end]
			if abs(target-sum) < abs(closestSum) {
				closestSum = target - sum
			}
			if sum < target {
				start++
			} else {
				end--
			}
		}
	}
	result = target - closestSum
	return
}

func abs(value int) (result int) {
	result = value
	if value < 0 {
		result = -1 * value
	}
	return
}