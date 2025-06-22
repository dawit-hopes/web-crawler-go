package main

import (
	"fmt"
	"io"
	"net/http"
)

func FetchURL(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	fmt.Println("Response status: ", resp.Status)

	body, sErr := io.ReadAll(resp.Body)

	if sErr != nil {
		return "", sErr
	}

	return string(body), nil

}
