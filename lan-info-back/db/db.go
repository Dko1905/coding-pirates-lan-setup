package db

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var _db *sqlx.DB

func init() {
	var err error
	_db, err = sqlx.Open("sqlite3", "./database.db")
	if err != nil {
		panic(err)
	}
}

func Begin(ctx context.Context) *sqlx.Tx {
	opts := sql.TxOptions{}
	return _db.MustBeginTx(ctx, &opts)
}
