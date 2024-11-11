package safesql

import (
	"brian/safesql/internal/raw"
	"context"
	"database/sql"
)

type compileTimeConstant string

type TrustedSQL struct {
	s string
}

func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{string(text)}
}

type DB struct {
	db *sql.DB
}

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*Rows, error) {
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

type Rows = sql.Rows
type Result = sql.Result

func init() {
	raw.TrustedSQLCtor =
		func(unsafe string) TrustedSQL {
			return TrustedSQL{unsafe}
		}
}
