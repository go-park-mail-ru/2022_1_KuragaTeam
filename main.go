package main

import (
	"log"
	"myapp/db"
	"myapp/handlers"
	"myapp/middleware"

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

	e.Use(middleware.CheckAuthorization(&connRedis))

	e.POST("/signup", handlers.CreateUserHandler(dbPool, &connRedis))
	e.POST("/login", handlers.LoginUserHandler(dbPool, &connRedis))
	e.DELETE("logout", handlers.LogoutHandler(&connRedis))
	e.GET("/", handlers.GetHomePageHandler(dbPool))
	e.Logger.Fatal(e.Start(":1323"))
}
