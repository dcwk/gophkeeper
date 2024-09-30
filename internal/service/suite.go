package service

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock_db "github.com/dcwk/gophkeeper/internal/infra/db/mocks"
	mock_repository "github.com/dcwk/gophkeeper/internal/repository/mocks"
)

type Suite struct {
	UserRepository *mock_repository.MockUserRepository
	TxManager      *mock_db.MockTxManager
}

func DefaultSuite(t *testing.T) *Suite {
	ctrl := gomock.NewController(t)
	s := Suite{}
	s.UserRepository = mock_repository.NewMockUserRepository(ctrl)
	s.TxManager = mock_db.NewMockTxManager(ctrl)

	return &s
}
