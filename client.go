package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func ProcessVinBDIOcr(testCase string, request VinBDIRequest) (VinBDIIDCardResponse, error) {
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

func GMOProcessIDCard(request GMOIDCardRequest) (*GMOIDCardResponse, error) {
	buffer := bytes.Buffer{}
	writer := multipart.NewWriter(&buffer)

	image1Field, err := writer.CreateFormFile("image1", frontImage)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(image1Field, bytes.NewBuffer(request.FrontPhotoContent))
	if err != nil {
		return nil, err
	}

	image2Field, err := writer.CreateFormFile("image2", rearImage)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(image2Field, bytes.NewBuffer(request.RearPhotoContent))
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("encode", "1")
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(http.MethodPost, GMORequestIdCardUrl, &buffer)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", writer.FormDataContentType())
	httpRequest.Header.Add("api-key", APIKeyGMO)
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
	var response GMOIDCardResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func GMOProcessFace(request GMOFaceRequest) (*GMOFaceResponse, error) {
	buffer := bytes.Buffer{}
	writer := multipart.NewWriter(&buffer)

	image1Field, err := writer.CreateFormFile("image1", faceImage)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(image1Field, bytes.NewBuffer(request.FacePhotoContent))
	if err != nil {
		return nil, err
	}

	image2Field, err := writer.CreateFormFile("image2", faceImage)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(image2Field, bytes.NewBuffer(request.IDPhotoROIContent))
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(http.MethodPost, GMORequestFaceUrl, &buffer)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", writer.FormDataContentType())
	httpRequest.Header.Add("api-key", APIKeyGMO)

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

	var response GMOFaceResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
