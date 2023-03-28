package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
	"os"
)

type DataBase struct {
	Client *sqlx.DB
}

func NewDataBase() (*DataBase, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	db, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return &DataBase{}, fmt.Errorf("could not connect to database: %w", err)
	}

	return &DataBase{Client: db}, nil

}

func (d *DataBase) Ping(ctx context.Context) error {
	return d.Client.DB.PingContext(ctx)
}
