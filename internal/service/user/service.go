package user

import "github.com/dcwk/gophkeeper/internal/repository"

type Service struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *Service {
	return &Service{userRepository: userRepository}
}
