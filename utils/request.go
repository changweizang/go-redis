package utils

type PhoneBody struct {
	Phone string `form:"phone" json:"phone"`
}

type Login struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
