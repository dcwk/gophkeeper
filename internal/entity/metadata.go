package entity

type Metadata struct {
	Id     int64
	Secret *Secret
	Key    string
	Value  string
}

func NewMetadata(key string, value string, secret *Secret) *Metadata {
	return &Metadata{
		Secret: secret,
		Key:    key,
		Value:  value,
	}
}
