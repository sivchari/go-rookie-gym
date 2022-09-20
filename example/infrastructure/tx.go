package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
)

type TxManager interface {
	BeginTransaction(fn func(tx TxManager) error) error
}

type Manager struct {
	db    *sql.DB
	tx    *sql.Tx
	inUse bool
}

func NewTxManager(db *sql.DB) TxManager {
	return newTxManager(db, nil, false)
}

func newTxManager(db *sql.DB, tx *sql.Tx, inUse bool) TxManager {
	return &Manager{
		db:    db,
		tx:    tx,
		inUse: inUse,
	}
}

func (m *Manager) BeginTransaction(fn func(tx TxManager) error) error {
	tx, err := m.DB().Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		p := recover()
		switch {
		case p != nil:
			if err2 := tx.Rollback(); err2 != nil {
				log.Printf("failed to Rollback: %v", err2)
			}
			panic(p)
		case err != nil:
			if err := tx.Rollback(); err != nil {
				log.Printf("failed to Rollback: %v", err)
			}
		default:
			if err := tx.Commit(); err != nil {
				log.Printf("failed to Commit: %v", err)
			}
		}
	}()
	err = fn(newTxManager(m.db, tx, true))
	return err
}

func (m *Manager) DB() *sql.DB {
	return m.db
}

func (m *Manager) Tx() *sql.Tx {
	return m.tx
}

func (m *Manager) InTransaction() bool {
	return m.inUse
}
