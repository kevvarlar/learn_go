package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	var user []User
	res, err := http.Get(url)
	if err != nil {
		return user, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}
