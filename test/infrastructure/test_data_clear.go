package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func TruncateTestData(ctx context.Context, dbpool *pgxpool.Pool) {
	_, truncateResultErr := dbpool.Exec(ctx, "TRUNCATE products RESTART IDENTITY")
	if truncateResultErr != nil {
		log.Error(truncateResultErr)
	} else {
		log.Info("Products table truncated")
	}
}
