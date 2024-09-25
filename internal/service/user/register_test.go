package user

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/service"
)

func TestRegisterUser(t *testing.T) {
	tests := []struct {
		Name           string
		User           entity.User
		PrepareMocks   func(s *service.Suite)
		ExpectedUserId int64
		ExpectedError  error
	}{
		{
			Name: "Success register user",
			User: entity.User{
				Login:    "test",
				Password: "test",
			},
			PrepareMocks: func(s *service.Suite) {
				s.UserRepository.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(int64(1), nil)
			},
			ExpectedUserId: 1,
			ExpectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			s := service.DefaultSuite(t)
			tt.PrepareMocks(s)
			userService := NewService(
				s.UserRepository,
				s.TxManager,
			)

			userId, err := userService.Register(context.Background(), tt.User)

			if tt.ExpectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.ExpectedUserId, userId)
			}
		})
	}
}
