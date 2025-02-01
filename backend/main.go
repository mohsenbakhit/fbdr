package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	app, err := setupFirebase()
	if err != nil {
		panic(fmt.Errorf("Error initializing Firebase: %v\n", err))
	}
	firebaseProjectID := os.Getenv("FIREBASE_PROJECT_ID")
	db, err := firestore.NewClient(ctx, firebaseProjectID)
	if err != nil {
		panic(fmt.Errorf("Error connecting to Firestore: %v\n", err))
	}
	r := setupRouter(db, app)
	err = r.Run(":8080")
	if err != nil {
		panic(fmt.Errorf("Error starting server: %v\n", err))
	}
}

func setupRouter(firestoreClient *firestore.Client, app *firebase.App) *gin.Engine {
	r := gin.Default()

	r.Use(corsMiddleware())

	api := r.Group("/api")
	{
		api.GET("/health", healthCheck)
		api.GET("/mlb-players", getMLBPlayerList)
		api.POST("/submit", submit(firestoreClient))
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
