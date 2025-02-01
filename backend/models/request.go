package models

type Request struct {
	Email           string   `json:"email"`
	FavoriteTeams   []string `json:"favoriteTeams"`
	FavoritePlayers []string `json:"favoritePlayers"`
}
