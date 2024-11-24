package db

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var BunDB *bun.DB

func SetBunDB(urlDB string) {
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(urlDB)))
	BunDB = bun.NewDB(sqlDB, pgdialect.New())
}
