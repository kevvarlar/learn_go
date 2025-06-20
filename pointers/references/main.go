package main

import (
	"strings"
)

func removeProfanity(message *string) {
	// ?
	profaneWords := map[string]string{
		"fubb" : "****",
		"shiz" : "****",
		"witch" : "*****",
	}
	for key, _ := range profaneWords {
		*message = strings.ReplaceAll(*message, key, profaneWords[key])
	}
}
