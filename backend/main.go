package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	app, err := setupFirebase()
	if err != nil {
		fmt.Errorf("Error initializing app: %v\n", err)
		return
	}
	db, err := app.Database(ctx)
	if err != nil {
		fmt.Errorf("Error initializing database: %v\n", err)
		return
	}
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		api.GET("/health", healthCheck)
		api.GET("/mlb-players", getMLBPlayerList)
	}

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupFirebase() (*firebase.App, error) {

	opt := option.WithCredentialsFile("firebase_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return app, nil

}
