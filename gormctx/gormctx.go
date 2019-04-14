package gormctx

import (
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
)

type DB struct {
	dialect string
	sqldb   *sql.DB

	// Options.
	logmode bool
}

func FromDB(dialect string, db *sql.DB) (*DB, error) {
	return &DB{dialect: dialect, sqldb: db}, nil
}

func (db *DB) WithContext(ctx context.Context) *gorm.DB {
	d, err := gorm.Open(db.dialect, &dbWithCtx{sqldb: db.sqldb, ctx: ctx})
	if err != nil {
		panic(err) // Should not happen as long as given db is valid.
	}
	if db.logmode {
		d.LogMode(db.logmode)
	}
	return d
}

func (db *DB) LogMode(mode bool) {
	db.logmode = mode
}

type dbWithCtx struct {
	sqldb *sql.DB
	ctx   context.Context
}

var _ gorm.SQLCommon = &dbWithCtx{}

func (db *dbWithCtx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.sqldb.ExecContext(db.ctx, query, args...)
}

func (db *dbWithCtx) Prepare(query string) (*sql.Stmt, error) {
	return db.sqldb.PrepareContext(db.ctx, query)
}

func (db *dbWithCtx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.sqldb.QueryContext(db.ctx, query, args...)
}

func (db *dbWithCtx) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.sqldb.QueryRowContext(db.ctx, query, args...)
}

func (db *dbWithCtx) Begin() (*sql.Tx, error) {
	return db.sqldb.BeginTx(db.ctx, nil)
}
