package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
	"github.com/dcwk/gophkeeper/internal/repository/user/converter"
	"github.com/dcwk/gophkeeper/internal/repository/user/model"
)

const (
	tableName = "\"user\""

	idColumn        = "id"
	loginColumn     = "login"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) CreateUser(ctx context.Context, user *entity.User) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(loginColumn, passwordColumn, createdAtColumn, updatedAtColumn).
		Values(user.Login, user.Password, sq.Expr("NOW()"), sq.Expr("NOW()")).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return id, nil
}

func (r *repo) GeUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	builder := sq.Select(idColumn, loginColumn, passwordColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{loginColumn: login})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GeUserByLogin",
		QueryRaw: query,
	}

	var user model.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
