package main

import (
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/mohsen/fbdr/models"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func getMLBPlayerList(c *gin.Context) {
	url := "https://statsapi.mlb.com/api/v1/sports/1/players?activeStatus=active"
	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "Internal Server Error",
			"message": "Error fetching MLB player list",
		})
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

func submit(f *firestore.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.Request
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		// Add document to Firestore
		_, _, err := f.Collection("users").Add(c, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save to Firestore",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User created successfully",
		})
	}
}
