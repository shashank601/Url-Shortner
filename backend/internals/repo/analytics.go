package repo

import "github.com/jackc/pgx/v5/pgxpool"

type AnalyticsRepo struct {
	DB *pgxpool.Pool
}
