package entity

type AuthPair struct {
	ID       int64
	Login    string
	Password string
	Secret   *Secret
}

func NewAuthPair(login string, password string, name string, metadata string) *AuthPair {
	return &AuthPair{
		Login:    login,
		Password: password,
		Secret:   NewSecret(AUTH_PAIR_TYPE, name, metadata),
	}
}
