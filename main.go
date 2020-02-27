package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	testCases, err := IOReadDir(DirectoryPath)
	if err != nil {
		panic(err)
	}

	//results := make([]VinBDIIDCardResponse, 0)
	results := make([]GMOTotalResponse, 0)
	total := len(testCases)

	for i, testCase := range testCases {
		err = preProcessImages(testCase)
		if err != nil {
			panic(err)
		}
		//frontImageBase64, err := base64File(testCase, frontImage)
		//if err != nil {
		//	panic(err)
		//}
		//faceImageBase64, err := base64File(testCase, faceImage)
		//if err != nil {
		//	panic(err)
		//}
		//rearImageBase64, err := base64File(testCase, rearImage)
		//if err != nil {
		//	panic(err)
		//}
		frontImage, err := ioutil.ReadFile(DirectoryPath + "/" + testCase + "/" + frontImage)
		if err != nil {
			panic(err)
		}
		rearImage, err := ioutil.ReadFile(DirectoryPath + "/" + testCase + "/" + rearImage)
		if err != nil {
			panic(err)
		}
		faceImage, err := ioutil.ReadFile(DirectoryPath + "/" + testCase + "/" + faceImage)
		if err != nil {
			panic(err)
		}

		//request := VinBDIRequest{
		//	FrontImageContentBase64: frontImageBase64,
		//	BackImageContentBase64:  rearImageBase64,
		//	FaceImageContentBase64:  faceImageBase64,
		//	IDBackType:              "upload",
		//	IDFrontType:             "upload",
		//}

		//response, err := ProcessOcr(testCase, request)

		requestID := GMOIDCardRequest{
			FrontPhotoContent: frontImage,
			RearPhotoContent:  rearImage,
		}

		var response2 *GMOFaceResponse
		response1, err := GMOProcessIDCard(requestID)
		fmt.Println("process: ", i, "/", total)
		if err != nil {
			fmt.Println("front request ---> false\n", testCase, err.Error())
			continue
		} else {
			imageRoiByte, err := base64.StdEncoding.DecodeString(response1.ImageROI)
			requestFace := GMOFaceRequest{
				FacePhotoContent:  faceImage,
				IDPhotoROIContent: imageRoiByte,
			}
			response2, err = GMOProcessFace(requestFace)
			if err != nil {
				fmt.Println("front request ---> false\n", testCase, err.Error())
				continue
			} else {
				fmt.Println("front request ---> success\n", testCase)
			}
		}
		response1.UserId = testCase
		result := MapIdCardAndFaceResponseToTotalResponse(*response1, *response2)
		results = append(results, result)
	}
	finish := time.Now()
	fmt.Println("Time has passed :", finish.Sub(start))

	writeResultsGMOToCsv(results)
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
	return err
}

func writeResultsVinBDIToCsv(results []VinBDIIDCardResponse) {
	f, err := os.Create("vinbdi_500.txt")
	if err != nil {
		return
	}
	defer f.Close()

	dataWriter := bufio.NewWriter(f)

	_, err = fmt.Fprintln(dataWriter, "user_id;id_type;id_number;full_name;dob;sex;nationality;home;address;expire_date;issue_date;issue_place;ethnicity;religion;face_score;face_matching")

	if err != nil {
		return
	}

	for _, result := range results {
		front := result.Data.FrontResponse
		back := result.Data.BackResponse
		face := result.Data.Face
		line := fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%v;%v;%s;%s;%s;%f;%t",
			result.UserId,
			front.Type,
			front.ID,
			front.Name,
			front.DOB,
			front.Sex,
			front.Nationality,
			front.Home,
			front.Address,
			front.DOE,
			back.IssueDate,
			back.IssuePlace,
			back.Ethnicity,
			back.Religion,
			face.MatchingScore,
			face.IsMatchingFace)
		_, err := fmt.Fprintln(dataWriter, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = dataWriter.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

}

func writeResultsGMOToCsv(results []GMOTotalResponse) {
	f, err := os.Create("GMO_500.txt")
	if err != nil {
		return
	}
	defer f.Close()

	dataWriter := bufio.NewWriter(f)

	_, err = fmt.Fprintln(dataWriter, "user_id;id_type;id_number;full_name;dob;sex;nationality;home;address;expire_date;issue_date;issue_place;ethnicity;religion;face_score;face_matching")

	if err != nil {
		return
	}

	for _, result := range results {
		var matchingScore float32
		var isMatchingFace bool
		isMatching, err := strconv.Atoi(result.FaceCompare)
		if err != nil {
			fmt.Println(err)
		}
		if isMatching == 1 {
			isMatchingFace = true
			matchingScore = 100
		} else {
			isMatchingFace = false
			matchingScore = 1
		}

		line := fmt.Sprintf("%s;%s;%s;%s;%s;%s;%s;%s;%s;%v;%v;%s;%s;%s;%f;%t",
			result.UserId,
			"",
			result.ID,
			result.Name,
			result.Birthday,
			result.Sex,
			"",
			result.HomeTown,
			result.Address,
			result.Expiry,
			result.IssueDate,
			result.IssueAt,
			"",
			"",
			matchingScore,
			isMatchingFace)
		_, err = fmt.Fprintln(dataWriter, line)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = dataWriter.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

}
