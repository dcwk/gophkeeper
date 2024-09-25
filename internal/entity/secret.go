package entity

const (
	AUTH_PAIR_TYPE = "AUTH_PAIR"
	TEXT_TYPE      = "TEXT"
	BINARY_TYPE    = "BINARY"
	PAYCARD_TYPE   = "PAYCARD"
)

type Secret struct {
	ID         int64
	User       User
	SecretType string
	Name       string
	MetaData   []*Metadata
}

func NewSecret(secretType string, name string, metadata string) *Secret {
	return &Secret{
		SecretType: secretType,
		Name:       name,
	}
}
