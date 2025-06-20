package main

import "fmt"
func getNameCounts(names []string) map[rune]map[string]int {
	// Your code here
	nameCounts := make(map[rune]map[string]int)

	for _, name := range names {
		first := []rune(name)[0]
		mm, ok := nameCounts[first]
		if !ok {
			mm = make(map[string]int)
			nameCounts[first] = mm
		}
		mm[name]++
	}
	fmt.Println(nameCounts)
	return nameCounts
}
