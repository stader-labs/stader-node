package net

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func MakePostRequest(requestUrl string, requestBody interface{}) (*http.Response, error) {
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func MakeGetRequest(requestUrl string, queryParameters interface{}) (*http.Response, error) {
	queryValues := url.Values{}
	queryParamsJSON, err := json.Marshal(queryParameters)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(queryParamsJSON, &queryValues)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", requestUrl, queryValues.Encode()), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
