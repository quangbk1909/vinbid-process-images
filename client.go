package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ProcessIDCardFront(request VinBDIIDCardRequest) (string, *VinBDIIDCardResponse, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return "", nil, err
	}
	writeBytes2Json("request", requestBytes)
	httpRequest, err := http.NewRequest(http.MethodPost, vinDBIURL, bytes.NewReader(requestBytes))
	if err != nil {
		return "", nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if httpResponse != nil && httpResponse.Body != nil {
		defer func() {
			_ = httpResponse.Body.Close()
		}()
	}

	if err != nil {
		return "", nil, err
	}

	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return "", nil, err
	}
	writeBytes2Json("all", bodyBytes)
	var response VinBDIIDCardResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return "", nil, err
	}

	return string(bodyBytes), &response, nil
}
