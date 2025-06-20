package main

func countConnections(groupSize int) int {
	count := 0
	for i := 0; i < groupSize; i++ {
		count += i
	}
	return count
}
