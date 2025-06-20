package main

func bulkSend(numMessages int) float64 {
	// ?
	if numMessages == 0 {
		return 0
	}
	cost := 0.
	fee := 0.
	for i := 0; i < numMessages; i++ {
		cost += 1 + fee
		fee += 0.01
	}
	return cost
}
