package request

type LoginStruct struct{
	Id  string `json:"id"`
	Password string `json:"Password"`
}

type RegisterStruct struct{
	Id  string `json:"id"`
	Password string `json:"Password"`
	NickName string `json:"NIckName"`
	ProfilePic string `json:"ProfilePic"`
}

type EditStruct struct{

	Password string `json:"Password"`
	NickName string `json:"NIckName"`
	ProfilePic string `json:"ProfilePic"`

}

type InfoStruct struct{
	Id  string `json:"id"`
	NickName string `json:"NIckName"`
	ProfilePic string `json:"ProfilePic"`
}