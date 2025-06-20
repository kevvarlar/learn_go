package main

func sum(nums ...int) int {
	// ?
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}
	return count
}
