package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ProcessOcr(testCase string, request VinBDIRequest) (VinBDIIDCardResponse, error) {
	//requestBytes, err := json.Marshal(request)
	//if err != nil {
	//	return VinBDIIDCardResponse{}, err
	//}
	//writeBytes2Json("request", testCase, requestBytes)
	//httpRequest, err := http.NewRequest(http.MethodPost, vinDBIURL, bytes.NewReader(requestBytes))
	//if err != nil {
	//	return VinBDIIDCardResponse{}, err
	//}
	//
	//httpRequest.Header.Set("Content-Type", "application/json")
	//httpClient := http.Client{}
	//httpResponse, err := httpClient.Do(httpRequest)
	//if httpResponse != nil && httpResponse.Body != nil {
	//	defer func() {
	//		_ = httpResponse.Body.Close()
	//	}()
	//}
	//if err != nil {
	//	return VinBDIIDCardResponse{}, err
	//}
	//bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	//if err != nil {
	//	return VinBDIIDCardResponse{}, err
	//}
	//writeBytes2Json("all", testCase, bodyBytes)
	bodyBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/all_output_vinbdi.json", DirectoryPath, testCase))
	var response VinBDIIDCardResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return VinBDIIDCardResponse{}, err
	}

	return response, nil
}

func writeBytes2Json(prefix, testcase string, input []byte) {
	fileName := fmt.Sprintf(`%v/%v/%v_output_vinbdi.json`, DirectoryPath, testcase, prefix)
	if err := ioutil.WriteFile(fileName, input, 0666); err != nil {
		fmt.Println("writeBytes2Json error: ", err.Error())
	}
}
