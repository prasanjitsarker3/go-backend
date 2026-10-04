package db

import (
	"backend-golang/config"
	"context"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" driver for database/sql
	"github.com/jmoiron/sqlx"
)

func GetConnectionString(cnf config.DBConfig) string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cnf.User, cnf.Password),
		Host:   fmt.Sprintf("%s:%d", cnf.Host, cnf.Port),
		Path:   cnf.Name,
	}
	q := u.Query()
	q.Set("sslmode", cnf.SSLMode)
	u.RawQuery = q.Encode()
	return u.String()
}

func NewConnection(ctx context.Context, cnf config.DBConfig) (*sqlx.DB, error) {
	dbCon, err := sqlx.ConnectContext(ctx, "pgx", GetConnectionString(cnf))
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
