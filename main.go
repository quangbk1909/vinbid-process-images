package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"time"
)

func base64File(filename string) (string, error) {
	fileContentBytes, err := ioutil.ReadFile("./images/" + TestCase + "/" + frontImage)
	if err != nil {
		fmt.Println("open file error: ", filename,  err.Error())
		return "", nil
	}

	fileContentBase64 := base64.StdEncoding.EncodeToString(fileContentBytes)
	return fileContentBase64, nil
}

func writeBytes2Json(prefix string, input []byte) {
	t := time.Now()
	_, month, date := t.Date()
	h := t.Hour()
	m := t.Minute()
	s := t.Second()
	fileName := fmt.Sprintf(`images/%v/%v_output_vinbdi_%v:%v-%v:%v:%v.json`, TestCase, prefix, month, date, h, m, s)
	if err := ioutil.WriteFile(fileName, input, 0666); err != nil {
		fmt.Println("writeBytes2Json error: ", err.Error())
	}

}

func main() {
	frontImageBase64, err := base64File(frontImage)
	if err != nil {
		panic(err)
	}

	faceImageBase64, err := base64File(faceImage)
	if err != nil {
		panic(err)
	}

	rearImageBase64, err := base64File(rearImage)
	if err != nil {
		panic(err)
	}

	vinBDIIDCardFrontRequest := VinBDIIDCardRequest{
		FrontImageContentBase64: frontImageBase64,
		BackImageContentBase64:  rearImageBase64,
		FaceImageContentBase64:  faceImageBase64,
		IDBackType:              "upload",
		IDFrontType:             "upload",
	}
	_, response, err := ProcessIDCardFront(vinBDIIDCardFrontRequest)
	if err != nil {
		fmt.Println("front request ---> false\n", err.Error())
	} else {
		fmt.Println("front request ---> success\n", response.Data.BackResponse.Errors, response.Data.Errors)
	}

}
