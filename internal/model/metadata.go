package model

type Metadata struct {
	Id     int64
	Secret Secret
	Key    string
	Value  string
}
