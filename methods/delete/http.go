package main

import (
	"errors"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

	req, err := http.NewRequest("DELETE", fullURL, nil)

	if err != nil {
		return err
	}

	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode > 299 {
		return errors.New("Unsuccessful delete")
	}
	return nil
}

