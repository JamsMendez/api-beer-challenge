package database

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"api-beer-challenge/settings"
)

func GetConnection(ctx context.Context, s *settings.Settings) (*sqlx.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		s.DBConfig.Host,
		s.DBConfig.Port,
		s.DBConfig.User,
		s.DBConfig.Password,
		s.DBConfig.Name,
	)

	return sqlx.ConnectContext(ctx, "postgres", connString)
}
