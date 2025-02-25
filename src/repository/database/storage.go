package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/M1RCLE/comments/src/config"
)

type Storage[T any] struct {
	*sqlx.DB
}

var DriverName = "postgres"

func NewStorage[T any](cfg config.Config) (*Storage[T], error) {
	db, err := sqlx.Connect(DriverName, cfg.DBConn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %v", err)
	}
	return &Storage[T]{db}, nil
}

func (d *Storage[T]) queryRowContext(ctx context.Context, query string, args ...interface{}) (T, error) {
	var record T

	tx, err := d.Beginx()
	if err != nil {
		return record, fmt.Errorf("transaction cannot start: %w", err)
	}

	err = tx.QueryRowxContext(ctx, query+" RETURNING *", args...).StructScan(&record)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return record, fmt.Errorf("transaction cannot execute query and rollback: %w, %w", err, errRollback)
		}
		return record, fmt.Errorf("transaction cannot execute query: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return record, fmt.Errorf("transaction cannot commit: %w", err)
	}

	return record, nil
}

func (d *Storage[T]) Insert(ctx context.Context, query string, args ...interface{}) (T, error) {
	return d.queryRowContext(ctx, query, args...)
}

func (d *Storage[T]) ExecuteUpdate(ctx context.Context, query string, args ...interface{}) (T, error) {
	return d.queryRowContext(ctx, query, args...)
}

func (d *Storage[T]) Delete(ctx context.Context, query string, args ...interface{}) error {
	tx, err := d.Beginx()
	if err != nil {
		return fmt.Errorf("transaction cannot start: %w", err)
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			return fmt.Errorf("transaction cannot execute DELETE and rollback: %w, %w", err, errRollback)
		}
		return fmt.Errorf("transaction cannot execute DELETE: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("transaction cannot commit: %w", err)
	}

	return nil
}
