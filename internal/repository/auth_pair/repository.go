package auth_pair

import (
	"context"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
)

const (
	tableName = "auth_pair"

	idColumn        = "id"
	secretIDColumn  = "secret_id"
	loginColumn     = "login"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.AuthPairRepository {
	return &repo{db: db}
}

func (r *repo) CreateAuthPair(ctx context.Context, authPair *entity.AuthPair) (int64, error) {
	//builder := sq.Insert(tableName).
	//	PlaceholderFormat(sq.Dollar).
	//	Columns(userIDColumn, typeColumn, nameColumn, createdAtColumn, updatedAtColumn).
	//	Values("1", secret.SecretType, secret.Name, sq.Expr("NOW()"), sq.Expr("NOW()")).
	//	Suffix("RETURNING id")
	//
	//query, args, err := builder.ToSql()
	//if err != nil {
	//	return 0, err
	//}
	//
	//q := db.Query{
	//	Name:     "secret_repository.CreateSecret",
	//	QueryRaw: query,
	//}
	//
	//var id int64
	//err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return 0, err
	//}
	//
	//return id, nil
	return 0, nil
}
