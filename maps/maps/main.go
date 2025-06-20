package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	// ?
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	userMap := map[string]user{}
	for i, name := range names {
		userMap[name] = user{name, phoneNumbers[i]}
	}
	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
