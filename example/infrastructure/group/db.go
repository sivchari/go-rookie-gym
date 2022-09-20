package groupdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/sivchari/go-rookie-gym/domain/group"
	"github.com/sivchari/go-rookie-gym/infrastructure"
)

type repository struct {
	m infrastructure.TxManager
}

var _ (group.Repository) = (*repository)(nil)

func NewDB(m infrastructure.TxManager) group.Repository {
	return &repository{
		m: m,
	}
}

func (r *repository) Transaction(txm infrastructure.TxManager) group.Repository {
	m, _ := txm.(*infrastructure.Manager)
	return NewDB(m)
}

func (r *repository) Group(ctx context.Context, group *group.Group) (int64, error) {
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
			"INSERT INTO `groups` (user_id, name) VALUES (?, ?)",
			group.UserID,
			group.Name,
		)
	} else {
		result, err = m.DB().ExecContext(
			ctx,
			"INSERT INTO `groups` (user_id, name) VALUES (?, ?)",
			group.UserID,
			group.Name,
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

func (r *repository) Groups(ctx context.Context, id int) ([]*group.Group, error) {
	m, ok := r.m.(*infrastructure.Manager)
	if !ok {
		return nil, errors.New("failed to cast to infrastructure.Manager")
	}
	var gs []*group.Group
	rows, err := m.DB().QueryContext(
		ctx,
		"SELECT * FROM `groups` WHERE user_id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var g group.Group
		if err := rows.Scan(&g.ID, &g.UserID, &g.Name); err != nil {
			return nil, err
		}
		gs = append(gs, &g)
	}

	return gs, nil
}
