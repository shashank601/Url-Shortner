package bootstrap

import (
	"github.com/shashank601/url-shortner/backend/internals/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB() *pgxpool.Pool {
	return db.NewDB()
}
