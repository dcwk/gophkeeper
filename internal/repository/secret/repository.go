package secret

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/dcwk/gophkeeper/internal/entity"
	"github.com/dcwk/gophkeeper/internal/infra/db"
	"github.com/dcwk/gophkeeper/internal/repository"
	"github.com/dcwk/gophkeeper/internal/repository/secret/converter"
	"github.com/dcwk/gophkeeper/internal/repository/secret/model"
)

const (
	tableName = "secret"

	idColumn        = "id"
	userIDColumn    = "user_id"
	typeColumn      = "type"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.SecretRepository {
	return &repo{db: db}
}

func (r *repo) CreateSecret(ctx context.Context, secret *entity.Secret) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIDColumn, typeColumn, nameColumn, createdAtColumn, updatedAtColumn).
		Values("1", secret.SecretType, secret.Name, sq.Expr("NOW()"), sq.Expr("NOW()")).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "secret_repository.CreateSecret",
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

func (r *repo) GetSecretsList(ctx context.Context) ([]*entity.Secret, error) {
	builder := sq.Select(idColumn, userIDColumn, typeColumn, nameColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "secret_repository.GetSecretsList",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	secrets := []*entity.Secret{}
	for rows.Next() {
		var secret model.Secret
		err := rows.Scan(
			&secret.ID,
			&secret.UserID,
			&secret.Type,
			&secret.Name,
			&secret.CreatedAt,
			&secret.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		secrets = append(secrets, converter.ToSecretFromRepo(&secret))
	}

	return secrets, nil
}
