package user

import "github.com/dcwk/gophkeeper/internal/repository"

type service struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *service {
	return &service{userRepository: userRepository}
}
