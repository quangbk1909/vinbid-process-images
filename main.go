package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	testCases, err := IOReadDir(DirectoryPath)
	if err != nil {
		panic(err)
	}
	for _, testCase := range testCases {
		err = preProcessImages(testCase)
		if err != nil {
			panic(err)
		}
		frontImageBase64, err := base64File(testCase, frontImage)
		if err != nil {
			panic(err)
		}
		faceImageBase64, err := base64File(testCase, faceImage)
		if err != nil {
			panic(err)
		}
		rearImageBase64, err := base64File(testCase, rearImage)
		if err != nil {
			panic(err)
		}

		request := VinBDIRequest{
			FrontImageContentBase64: frontImageBase64,
			BackImageContentBase64:  rearImageBase64,
			FaceImageContentBase64:  faceImageBase64,
			IDBackType:              "upload",
			IDFrontType:             "upload",
		}

		response, err := ProcessOcr(testCase, request)
		if err != nil {
			fmt.Println("front request ---> false\n", testCase, err.Error())
		} else {
			fmt.Println("front request ---> success\n", testCase, response.Data.BackResponse.Errors, response.Data.Errors)
		}
	}
	finish := time.Now()
	fmt.Println("Time has passed :", finish.Sub(start))
}

func IOReadDir(parentDir string) ([]string, error) {
	var childDir []string
	fileInfo, err := ioutil.ReadDir(parentDir)
	if err != nil {
		return childDir, err
	}

	for _, file := range fileInfo {
		if file.IsDir() {
			childDir = append(childDir, file.Name())
		}
	}
	return childDir, nil
}

func base64File(testcase, filename string) (string, error) {
	fileContentBytes, err := ioutil.ReadFile(DirectoryPath + "/" + testcase + "/" + filename)
	if err != nil {
		fmt.Println("open file error: ", filename, err.Error())
		return "", nil
	}
	fileContentBase64 := base64.StdEncoding.EncodeToString(fileContentBytes)
	return fileContentBase64, nil
}

func preProcessImages(testCase string) (err error) {
	files, err := ioutil.ReadDir(DirectoryPath + "/" + testCase)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir() {
			if strings.Contains(file.Name(), "front") && file.Name() != frontImage {
				originalPath := DirectoryPath + "/" + testCase + "/" + file.Name()
				newPath := DirectoryPath + "/" + testCase + "/" + frontImage
				err = os.Rename(originalPath, newPath)
				if err != nil {
					return err
				}
			}
			if strings.Contains(file.Name(), "rear") && file.Name() != rearImage {
				originalPath := DirectoryPath + "/" + testCase + "/" + file.Name()
				newPath := DirectoryPath + "/" + testCase + "/" + rearImage
				err = os.Rename(originalPath, newPath)
				if err != nil {
					return err
				}
			}
			if strings.Contains(file.Name(), "face") && file.Name() != faceImage {
				originalPath := DirectoryPath + "/" + testCase + "/" + file.Name()
				newPath := DirectoryPath + "/" + testCase + "/" + faceImage
				err = os.Rename(originalPath, newPath)
				if err != nil {
					return err
				}
			}
		}
	}
	return
}
