package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sivchari/go-rookie-gym/domain/user"
	"github.com/sivchari/go-rookie-gym/infrastructure"
)

type repository struct {
	m infrastructure.TxManager
}

var _ (user.Repository) = (*repository)(nil)

func NewDB(m infrastructure.TxManager) user.Repository {
	return &repository{
		m: m,
	}
}

func (r *repository) Transaction(txm infrastructure.TxManager) user.Repository {
	m, _ := txm.(*infrastructure.Manager)
	return NewDB(m)
}

func (r *repository) User(ctx context.Context, user *user.User) (int64, error) {
	m, ok := r.m.(*infrastructure.Manager)
	if !ok {
		return 0, errors.New("failed to cast to infrastructure.Manager")
	}
	var (
		result sql.Result
		err    error
	)
	if m.InTransaction() {
		result, err = m.Tx().ExecContext(
			ctx,
			"INSERT INTO `users` (name) VALUES (?)",
			user.Name,
		)
	} else {
		result, err = m.DB().ExecContext(
			ctx,
			"INSERT INTO `users` (name) VALUES (?)",
			user.Name,
		)
	}
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
