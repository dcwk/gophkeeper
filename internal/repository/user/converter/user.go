package converter

import (
	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/repository/user/model"
)

func ToUserFromRepo(user *model.User) *entity.User {
	return &entity.User{
		ID:       user.Id,
		Login:    user.Login,
		Password: user.Password,
	}
}
