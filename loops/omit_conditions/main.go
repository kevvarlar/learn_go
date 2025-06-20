package main

func maxMessages(thresh int) int {
	cost := 0
	fee := 0
	count := 0
	for i := 0; ;i++ {
		if cost > thresh {
			count--
			break;
		}
		cost += 100 + fee
		fee++
		count++
	}
	return count
}
