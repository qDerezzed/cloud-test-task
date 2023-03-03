package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type PostgresDB struct {
	databaseURL string
	dbPool      *pgxpool.Pool
}

func New(cfg Config) *PostgresDB {
	return &PostgresDB{
		databaseURL: fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode),
	}
}

func (db *PostgresDB) Open() error {
	dbpool, err := pgxpool.Connect(context.Background(), db.databaseURL)
	db.dbPool = dbpool
	return err
}

func (db *PostgresDB) Close() {
	db.dbPool.Close()
}
