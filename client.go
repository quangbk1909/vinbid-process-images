package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func ProcessOcr(testCase string, request VinBDIRequest) (*VinBDIIDCardResponse, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	writeBytes2Json("request", testCase, requestBytes)
	httpRequest, err := http.NewRequest(http.MethodPost, vinDBIURL, bytes.NewReader(requestBytes))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	writeBytes2Json("all", testCase, bodyBytes)
	var response VinBDIIDCardResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func writeBytes2Json(prefix, testcase string, input []byte) {
	t := time.Now()
	_, month, date := t.Date()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	fileName := fmt.Sprintf(`%v/%v/%v_output_vinbdi_%v:%v-%v:%v:%v.json`, DirectoryPath, testcase, prefix, month, date, h, m, s)
	if err := ioutil.WriteFile(fileName, input, 0666); err != nil {
		fmt.Println("writeBytes2Json error: ", err.Error())
	}
}
