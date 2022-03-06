package main

import (
	"log"
	"myapp/db"
	"myapp/handlers"

	_ "github.com/jackc/pgx/v4"
	"github.com/labstack/echo"
)

func main() {
	dbPool, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	connRedis, err := db.ConnectRedis()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = connRedis.Close()
		if err != nil {
			log.Fatal(err)
		}
		dbPool.Close()
	}()

	e := echo.New()

	e.POST("/signup", handlers.CreateUserHandler(dbPool, &connRedis))
	e.POST("/signin", handlers.LoginUserHandler(dbPool, &connRedis))
	e.GET("/", handlers.GetHomePageHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
