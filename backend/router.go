package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
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
