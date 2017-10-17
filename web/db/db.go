package db

import (
	"github.com/jackc/pgx"
	"os"
)

var dbPool *pgx.ConnPool

func init() {
	var err error
	connPoolConfig := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			User:     "postgres",
			Password: "",
			Database: "cc_fileshare",
		},
		MaxConnections: 5,
	}
	dbPool, err = pgx.NewConnPool(connPoolConfig)
	if err != nil {
		os.Exit(1)
	}
}

func GetPool() *pgx.ConnPool {
	return dbPool
}
