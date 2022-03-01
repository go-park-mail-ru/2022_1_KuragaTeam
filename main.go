package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"log"
	"myapp/handlers"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "net"
	password = "pass"
	dbname   = "netflix"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()

	e.POST("/signup", handlers.CreateUserHandler(db))
	e.POST("/signin", handlers.LoginUserHandler(db))
	e.GET("/", handlers.GetHomePageHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
