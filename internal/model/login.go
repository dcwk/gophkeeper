package model

type Login struct {
	Id       int64
	Secret   Secret
	Login    string
	Password string
}
