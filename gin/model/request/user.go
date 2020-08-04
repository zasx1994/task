package request

type LoginStruct struct{
	Id  string
	Password string
}

type RegisterStruct struct{
	Id  string
	Password string
	NickName string
	ProfilePic string
}

type EditStruct struct{
	Id string
	Password string
	NickName string
	ProfilePic string

}

type InfoStruct struct{
	Id  string
	NickName string
	ProfilePic string
}