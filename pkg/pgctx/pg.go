package pgctx

import (
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgconn"
)

// UniqueViolation check unique constraint error
func UniqueViolation(err error, name string) bool {
	var pqErr *pgconn.PgError
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return strings.Contains(pqErr.Message, name)
	}
	return false
}

// ForeignKeyViolation check foreign key constraint error
func ForeignKeyViolation(err error, name string) bool {
	var pqErr *pgconn.PgError
	if errors.As(err, &pqErr) && pqErr.Code == "23503" {
		return strings.Contains(pqErr.Message, name)
	}
	return false
}
