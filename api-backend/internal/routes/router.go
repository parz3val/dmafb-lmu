package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	// "log"
	"net/http"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func Router(db *mongo.Client) *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.GET("/", HealthCheck)

	// v1 := e.Group("/api/v1")
	return e
}
