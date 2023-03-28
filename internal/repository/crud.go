package repository

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFoundEntity     = errors.New("not found entity")
	ErrNoneEntityDeleted  = errors.New("none entity rows deleted in database")
	ErrNoneEntityInserted = errors.New("none entity rows inserted in database")
	ErrNoneEntityUpdated  = errors.New("none entity rows updated in database")
	ErrNoColums           = errors.New("not found columns to update")
)

func (r *repository) EqualsNotDeletedItems(err error) bool {
	return errors.Is(err, ErrNoneEntityDeleted)
}

func (r *repository) insertEntity(ctx context.Context, query string, fields ...any) (id uint64, err error) {
	err = r.db.GetContext(ctx, &id, query, fields...)
	if err != nil {
		return id, err
	}

	return id, err
}

func (r *repository) updateEntity(ctx context.Context, query string, id uint64, fields ...any) error {
	var result sql.Result

	// implement preppend
	fields = append([]any{id}, fields...)

	result, err := r.db.ExecContext(ctx, query, fields...)
	if err != nil {
		return err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsUpdated == 0 {
		return ErrNoneEntityUpdated
	}

	return err
}

func (r *repository) deleteEntity(ctx context.Context, query string, id uint64) error {
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsDeleted == 0 {
		return ErrNoneEntityDeleted
	}

	return err
}
