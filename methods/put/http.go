package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
	var user User
	jsonData, err := json.Marshal(data)
	if err != nil {
		return user, err
	}

	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))

	if err != nil {
		return user, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)

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

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id
	var user User
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return user, err
	}
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)
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

