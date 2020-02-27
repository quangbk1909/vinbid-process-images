package main

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

type GMOIDCardResponse struct {
	UserId       string `json:"-"`
	Address      string `json:"address"`
	Birthday     string `json:"birthday"`
	Expiry       string `json:"expiry"`
	FirstName    string `json:"first_name"`
	FrontFlg     int    `json:"front_flg"`
	ID           string `json:"id"`
	ImageROI     string `json:"image_roi"`
	HomeTown     string `json:"home_town"`
	IssueAt      string `json:"issue_at"`
	IssueDate    string `json:"issue_date"`
	LastName     string `json:"last_name"`
	LogicCheck   string `json:"logiccheck"`
	LogicMessage string `json:"logicmessage"`
	Name         string `json:"name"`
	ResultCode   int    `json:"result_code"`
	Sex          string `json:"sex"`
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

type GMOTotalResponse struct {
	UserId       string `json:"-"`
	Address      string `json:"address"`
	Birthday     string `json:"birthday"`
	Expiry       string `json:"expiry"`
	FirstName    string `json:"first_name"`
	FrontFlg     int    `json:"front_flg"`
	ID           string `json:"id"`
	ImageROI     string `json:"image_roi"`
	HomeTown     string `json:"home_town"`
	IssueAt      string `json:"issue_at"`
	IssueDate    string `json:"issue_date"`
	LastName     string `json:"last_name"`
	LogicCheck   string `json:"logiccheck"`
	LogicMessage string `json:"logicmessage"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	FaceCompare  string `json:"face_compare"`
	Message      string `json:"message"`
}

func MapIdCardAndFaceResponseToTotalResponse(idCardResponse GMOIDCardResponse, faceResponse GMOFaceResponse) GMOTotalResponse {
	return GMOTotalResponse{
		UserId:      idCardResponse.UserId,
		Address:     idCardResponse.Address,
		Birthday:    idCardResponse.Birthday,
		Expiry:      idCardResponse.Expiry,
		FirstName:   idCardResponse.FirstName,
		FrontFlg:    idCardResponse.FrontFlg,
		ID:          idCardResponse.ID,
		ImageROI:    idCardResponse.ImageROI,
		HomeTown:    idCardResponse.HomeTown,
		IssueAt:     idCardResponse.IssueAt,
		IssueDate:   idCardResponse.IssueDate,
		LastName:    idCardResponse.LastName,
		LogicCheck:  idCardResponse.LogicCheck,
		Name:        idCardResponse.Name,
		Sex:         idCardResponse.Sex,
		FaceCompare: faceResponse.FaceCompare,
		Message:     faceResponse.Message,
	}
}
