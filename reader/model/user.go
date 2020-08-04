package model

type User struct {
	Id string `json:"Id"`
	NickName string `json:"NickName"`
	PassWord string `json:"Password"`
	ProfilePic string `json:"ProfilePic"`
}
