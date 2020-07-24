package model

type User struct {
	Id string `json:"id"`
	NickName string `json:"nickname"`
	PassWord string `json:"password"`
	ProfilePic string `json:"profilepic"`
}
