package db

import (
	"database/sql"

	"entgo.io/ent/dialect"
	entSQL "entgo.io/ent/dialect/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"api-go/ent"
)

var EntDB *ent.Client
var BunDB *bun.DB

func SetDB(urlDB string) {
	connDB, err := sql.Open("pgx", urlDB)
	if err != nil {
		panic("Failed to open database connection")
	}

	err = connDB.Ping()
	if err != nil {
		panic("Failed to ping database")
	}

	sqlDB := entSQL.OpenDB(dialect.Postgres, connDB)
	EntDB = ent.NewClient(ent.Driver(sqlDB))
}

func SetBunDB(urlDB string) {
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(urlDB)))
	BunDB = bun.NewDB(sqlDB, pgdialect.New())
}
