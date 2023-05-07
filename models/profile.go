package models

type Profile struct {
	Id             string
	UserId         string
	FullName       string `json:"fullname"`
	AvatarFileName string `json:"avatar"`
}
