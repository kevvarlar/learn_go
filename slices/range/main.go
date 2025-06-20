package main

import "slices"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	// ?
	for i, word := range msg {
		if slices.Contains(badWords, word) {
				return i
			}
	}
	return -1
}
