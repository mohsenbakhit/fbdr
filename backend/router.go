package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohsenbakhit/fbdr/models"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

// getMLBPlayerList fetches a list of active MLB players from statsapi.mlb.com
func GetMLBPlayerList(c *gin.Context) {
	url := "https://statsapi.mlb.com/api/v1/sports/1/players?activeStatus=active"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "Internal Server Error",
			"message": "Error fetching MLB player list",
		})
		return
	}
	defer resp.Body.Close()

	var result struct {
		Players []struct {
			FullName string `json:"fullName"`
		} `json:"people"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(500, gin.H{
			"status":  "Internal Server Error",
			"message": "Error decoding MLB player list",
		})
		return
	}

	var playerNames []string
	for _, player := range result.Players {
		playerNames = append(playerNames, player.FullName)
	}

	c.JSON(200, gin.H{
		"status":  "ok",
		"players": playerNames,
	})
}

// submit takes the body of the POST request and calls the Gemini function
func Submit(c *gin.Context) {
	var user models.Request
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": "Invalid request body",
		})
		return
	}

	_, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "Internal Server Error",
			"message": "Error inserting user",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
