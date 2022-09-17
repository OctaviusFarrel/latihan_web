package responses

import "octaviusfarrel.dev/latihan_web/models"

type (
	AllPlayers struct {
		BaseResponse
		Players []models.PlayerModel `json:"players"`
	}

	Player struct {
		BaseResponse
		Player models.PlayerModel `json:"player"`
	}
)
