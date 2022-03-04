package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"),
		os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	dbPool, err := pgxpool.Connect(context.Background(), psqlInfo)

	if err != nil {
		return nil, err
	}

	return dbPool, nil
}
