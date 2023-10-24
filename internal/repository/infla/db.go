package infla

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var dsn = "postgres://postgres:@localhost:5432/postgres?sslmode=disable"
var db = bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn))), pgdialect.New())

func NewDB() *bun.DB {
	return db
}
