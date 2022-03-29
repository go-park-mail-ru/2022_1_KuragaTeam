package composites

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/joho/godotenv"
)

type PostgresDBComposite struct {
	db *sql.DB
}

func NewPostgresDBComposite() (*PostgresDBComposite, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))

	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping() // вот тут будет первое подключение к базе
	if err != nil {
		return nil, err
	}

	return &PostgresDBComposite{db: db}, nil
}
