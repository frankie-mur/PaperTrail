package main

import (
	"net/http"
	"os"

	"github.com/frankie-mur/PaperTrail/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Load up env vars
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env vars")
	}

	connectionUri := os.Getenv("MONGO_URI")
	mongoDb, err := db.NewMongoDB(connectionUri, "admin", "")
	if err != nil {
		panic("failed to connect to db")
	}

	// //Create Our Services
	// articleService := article.NewArticleService(mongoDb)

	//Init our routes
	e.GET("/health", healthHandler(mongoDb))
	e.Logger.Fatal(e.Start(":8080"))
}

func healthHandler(db *db.MongoDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Example health check: Ping the database
		err := db.Ping(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"status": "unhealthy", "error": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	}
}
