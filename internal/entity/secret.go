package entity

type Secret struct {
	Id         int64
	User       User
	SecretType string
	Name       string
	MetaData   []Metadata
}
