package main

import (
	"log"

	"go-ws/task-app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"
)

func init() {
	err := gotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	apiGroupV1 := e.Group("/api/v1")
	routes.Route(apiGroupV1)

	log.Println("Server is running at port 8080")
	log.Fatal(e.Start(":8080"))
}
