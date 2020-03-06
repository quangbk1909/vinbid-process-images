package main

import (
	"strconv"
	"time"
)

type VinBDIRequest struct {
	FrontImageContentBase64 string `json:"id_front_img"`
	BackImageContentBase64  string `json:"id_back_img"`
	FaceImageContentBase64  string `json:"face_img"`
	IDBackType              string `json:"id_back_type"`
	IDFrontType             string `json:"id_front_type"`
}

type VinBDIError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type VinBDICommonImageDataResponse struct {
	ID                string        `json:"id"`
	IDProb            string        `json:"id_prob"`
	IDStatus          string        `json:"id_status"`
	IDCharacterStatus string        `json:"id_character_status"`
	Address           string        `json:"address"`
	AddressStatus     string        `json:"address_status"`
	AddressProb       string        `json:"address_prob"`
	Name              string        `json:"name"`
	NameStatus        string        `json:"name_status"`
	NameProb          string        `json:"name_prob"`
	Nationality       string        `json:"nationality"`
	NationalityProb   string        `json:"nationality_prob"`
	NationalityStatus string        `json:"nationality_status"`
	DOB               interface{}   `json:"dob"`
	DOBProb           string        `json:"dob_prob"`
	DOBStatus         string        `json:"dob_status"`
	Sex               string        `json:"sex"`
	SexProb           string        `json:"sex_prob"`
	SexStatus         string        `json:"sex_status"`
	Ethnicity         string        `json:"ethnicity"`
	EthnicityProb     string        `json:"ethnicity_prob"`
	EthnicityStatus   string        `json:"ethnicity_status"`
	Home              string        `json:"home"`
	HomeProb          string        `json:"home_prob"`
	HomeStatus        string        `json:"home_status"`
	DOE               interface{}   `json:"doe"`
	DOEPrb            string        `json:"doe_prb"`
	DOEStatus         string        `json:"doe_status"`
	IssueDate         interface{}   `json:"issue_date"`
	IssueDateProb     string        `json:"issue_date_prob"`
	IssueDateStatus   string        `json:"issue_date_status"`
	IssuePlace        string        `json:"issue_loc"`
	IssuePlaceProb    string        `json:"issue_loc_prob"`
	IssuePlaceStatus  string        `json:"issue_loc_status"`
	Religion          string        `json:"religion"`
	ReligionProb      string        `json:"religion_prob"`
	ReligionStatus    string        `json:"religion_status"`
	IsSmall           bool          `json:"isSmall"`
	IsCropper         bool          `json:"isCropper"`
	RotateDegree      float64       `json:"rotate_degree"`
	Errors            []VinBDIError `json:"errors"`
	Type              string        `json:"type"`
}

type VinBDIIDCardResponse struct {
	UserId  string `json:"-"`
	Message string `json:"message"`
	Code    int64  `json:"Code"`
	Data    struct {
		FrontResponse VinBDICommonImageDataResponse `json:"id_front"`
		BackResponse  VinBDICommonImageDataResponse `json:"id_back"`
		IDRecapture   struct {
			IDFront int64 `json:"id_front"`
			IDBack  int64 `json:"id_back"`
		} `json:"id_recapture"`

		Face struct {
			MatchingScore  float64       `json:"matching_score"`
			IsMatchingFace bool          `json:"is_matching_face"`
			IsRecapture    string        `json:"is_recapture"`
			RecaptureScore bool          `json:"recapture_score"`
			Errors         []VinBDIError `json:"errors"`
		} `json:"face"`

		Errors []VinBDIError `json:"errors"`
	} `json:"data"`
}

type GMOIDCardRequest struct {
	FrontPhotoContent []byte
	RearPhotoContent  []byte
}

//type GMOIDCardResponse struct {
//	UserId       string `json:"-"`
//	Address      string `json:"address"`
//	Birthday     string `json:"birthday"`
//	Expiry       string `json:"expiry"`
//	FirstName    string `json:"first_name"`
//	FrontFlg     int    `json:"front_flg"`
//	ID           string `json:"id"`
//	ImageROI     string `json:"image_roi"`
//	HomeTown     string `json:"home_town"`
//	IssueAt      string `json:"issue_at"`
//	IssueDate    string `json:"issue_date"`
//	LastName     string `json:"last_name"`
//	LogicCheck   string `json:"logiccheck"`
//	LogicMessage string `json:"logicmessage"`
//	Name         string `json:"name"`
//	ResultCode   int    `json:"result_code"`
//	Sex          string `json:"sex"`
//}

type GMOIDCardResponse struct {
	UserId       string      `json:"-"`
	Address      interface{} `json:"address"`
	Birthday     interface{} `json:"birthday"`
	CornerCheck  string      `json:"corner_check"`
	Expiry       interface{} `json:"expiry"`
	FrontFlg     int         `json:"front_flg"`
	HomeTown     interface{} `json:"home_town"`
	ID           interface{} `json:"id"`
	ImageROI     string      `json:"image_roi"`
	IssueAt      interface{} `json:"issue_at"`
	IssueDate    interface{} `json:"issue_date"`
	LogicCheck   string      `json:"logiccheck"`
	LogicMessage string      `json:"logicmessage"`
	Name         interface{} `json:"name"`
	ResultCode   int         `json:"result_code"`
	Sex          interface{} `json:"sex"`
	Type         string      `json:"type"`
}

type GMOFaceRequest struct {
	FacePhotoContent  []byte
	IDPhotoROIContent []byte
}

type GMOFaceResponse struct {
	FaceCompare string `json:"face_compare"`
	Message     string `json:"message"`
	ResultCode  int    `json:"result_code"`
}

//type GMOTotalResponse struct {
//	UserId       string `json:"-"`
//	Address      string `json:"address"`
//	Birthday     string `json:"birthday"`
//	Expiry       string `json:"expiry"`
//	FirstName    string `json:"first_name"`
//	FrontFlg     int    `json:"front_flg"`
//	ID           string `json:"id"`
//	ImageROI     string `json:"image_roi"`
//	HomeTown     string `json:"home_town"`
//	IssueAt      string `json:"issue_at"`
//	IssueDate    string `json:"issue_date"`
//	LastName     string `json:"last_name"`
//	LogicCheck   string `json:"logiccheck"`
//	LogicMessage string `json:"logicmessage"`
//	Name         string `json:"name"`
//	Sex          string `json:"sex"`
//	FaceCompare  string `json:"face_compare"`
//	Message      string `json:"message"`
//}

type GMOOCRResult struct {
	UserId                string     `json:"-"`
	FullName              string     `json:"full_name"`
	FullNameScore         float32    `json:"full_name_score"`
	DOBVal                *time.Time `json:"dob_val"`
	DOB                   string     `json:"dob"`
	DOBScore              float32    `json:"dob_score"`
	Gender                string     `json:"gender"`
	GenderScore           float32    `json:"gender_score"`
	IdentityType          string     `json:"identity_type"`
	IdentityNumber        string     `json:"identity_number"`
	IdentityNumberScore   float32    `json:"identity_number_score"`
	Nationality           string     `json:"nationality"`
	NationalityScore      float32    `json:"nationality_score"`
	PermanentAddress      string     `json:"permanent_address"`
	PermanentAddressScore float32    `json:"permanent_address_score"`
	HometownAddress       string     `json:"hometown_address"`
	HometownAddressScore  float32    `json:"hometown_address_score"`
	IssueDateVal          *time.Time `json:"issue_date_val"`
	IssueDate             string     `json:"issue_date"`
	IssueDateScore        float32    `json:"issue_date_score"`
	IssuePlace            string     `json:"issue_place"`
	IssuePlaceScore       float32    `json:"issue_place_score"`
	ExpireDateVal         *time.Time `json:"expire_date_val"`
	ExpireDate            string     `json:"expire_date"`
	ExpireDateScore       float32    `json:"expire_date_score"`
	FaceCompare           string     `json:"face_compare"`
	FaceMessage           string     `json:"face_message"`
}

//func MapIdCardAndFaceResponseToTotalResponse(idCardResponse GMOIDCardResponse, faceResponse GMOFaceResponse) GMOTotalResponse {
//	return GMOTotalResponse{
//		UserId:      idCardResponse.UserId,
//		Address:     idCardResponse.Address,
//		Birthday:    idCardResponse.Birthday,
//		Expiry:      idCardResponse.Expiry,
//		FirstName:   idCardResponse.FirstName,
//		FrontFlg:    idCardResponse.FrontFlg,
//		ID:          idCardResponse.ID,
//		ImageROI:    idCardResponse.ImageROI,
//		HomeTown:    idCardResponse.HomeTown,
//		IssueAt:     idCardResponse.IssueAt,
//		IssueDate:   idCardResponse.IssueDate,
//		LastName:    idCardResponse.LastName,
//		LogicCheck:  idCardResponse.LogicCheck,
//		Name:        idCardResponse.Name,
//		Sex:         idCardResponse.Sex,
//		FaceCompare: faceResponse.FaceCompare,
//		Message:     faceResponse.Message,
//	}
//}

func extractOCRResult(input *GMOIDCardResponse, identityType string) (GMOOCRResult, []error) {
	output := GMOOCRResult{
		IdentityType: identityType,
	}

	output.FullName, output.FullNameScore = extractGMOData(input.Name)
	output.Gender, output.GenderScore = extractGMOData(input.Sex)
	output.IdentityNumber, output.IdentityNumberScore = extractGMOData(input.ID)
	output.PermanentAddress, output.PermanentAddressScore = extractGMOData(input.Address)
	output.HometownAddress, output.HometownAddressScore = extractGMOData(input.HomeTown)
	output.IssuePlace, output.IssuePlaceScore = extractGMOData(input.IssueAt)

	inputDateFormat := "2/1/2006"
	outputDateFormat := "2006-01-02"
	errs := make([]error, 0)

	birthdayText, birthdayScore := extractGMOData(input.Birthday)
	if birthdayText != "" {
		birthdayVal, err := time.Parse(inputDateFormat, birthdayText)
		if err == nil {
			output.DOBVal = &birthdayVal
			output.DOB = birthdayVal.Format(outputDateFormat)
			output.DOBScore = birthdayScore
		} else {
			errs = append(errs, err)
		}
	}

	issueDateText, issueDateScore := extractGMOData(input.IssueDate)
	if issueDateText != "" {
		issueDateVal, err := time.Parse(inputDateFormat, issueDateText)
		if err == nil {
			output.IssueDateVal = &issueDateVal
			output.IssueDate = issueDateVal.Format(outputDateFormat)
			output.IssueDateScore = issueDateScore
		} else {
			errs = append(errs, err)
		}
	}

	expiryText, expiryScore := extractGMOData(input.Expiry)
	if expiryText != "" {
		expiryVal, err := time.Parse(inputDateFormat, expiryText)
		if err == nil {
			output.ExpireDateVal = &expiryVal
			output.ExpireDate = expiryVal.Format(outputDateFormat)
			output.ExpireDateScore = expiryScore
		} else {
			errs = append(errs, err)
		}
	}

	return output, errs
}

func extractGMOData(input interface{}) (text string, score float32) {
	defer func() {
		score = score * 100
	}()

	if input == nil {
		return
	}

	var ok bool
	if text, ok = input.(string); ok {
		return
	}

	if val, ok := input.(map[string]interface{}); ok {
		if val["text"] != nil {
			text, _ = val["text"].(string)
		}

		if val["score"] == nil {
			return
		}

		score, ok = val["score"].(float32)
		if ok {
			return
		}

		score64, ok := val["score"].(float64)
		if ok {
			score = float32(score64)
			return
		}

		scoreInt, ok := val["score"].(int)
		if ok {
			score = float32(scoreInt)
			return
		}

		scoreStr, ok := val["score"].(string)
		if ok {
			floatVal, err := strconv.ParseFloat(scoreStr, 64)
			if err != nil {
				return
			}

			score = float32(floatVal)
			return
		}
	}

	return
}
