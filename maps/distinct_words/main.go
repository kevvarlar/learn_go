package main

import (
	"strings"
	"fmt"
)
func countDistinctWords(messages []string) int {
	count := 0
	freqCounter := make(map[string]int)
	for _, message := range messages {
		currentString := ""
		currentString = ""
		for i, c := range message {
			if string(c) == " "  || i == len(message) - 1 {
				if i == len(message) - 1 {
					currentString += string(c)
				}
				currentString = strings.ToLower(currentString)
				fmt.Println("Outside: " + currentString)
				if _, ok := freqCounter[currentString]; !ok {
					count += 1
					fmt.Println("Inside: " + currentString)
					freqCounter[currentString]++
				}
				currentString = ""
			} else {
				currentString += string(c)
			}
		}
	}
	return count
}
