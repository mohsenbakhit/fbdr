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

func getMLBPlayerList() ([]string, error) {
	url := "https://statsapi.mlb.com/api/v1/sports/1/players?activeStatus=active"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Players []struct {
			FullName string `json:"fullName"`
		} `json:"people"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var playerNames []string
	for _, player := range result.Players {
		playerNames = append(playerNames, player.FullName)
	}

	return playerNames, nil
}
