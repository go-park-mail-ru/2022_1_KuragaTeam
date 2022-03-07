package main

import (
	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
	"log"
	"myapp/db"
	"myapp/handlers"

	_ "github.com/jackc/pgx/v4"
)

// @title Movie Space API
// @version 1.0
// @description This is API server for Movie Space website.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api/v1/
// @schemes http
func main() {
	dbPool, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer dbPool.Close()

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/api/v1/signup", handlers.CreateUserHandler(dbPool))
	e.POST("/api/v1/signin", handlers.LoginUserHandler(dbPool))
	e.GET("/api/v1/", handlers.GetHomePageHandler())
	e.Logger.Fatal(e.Start(":1323"))
}
