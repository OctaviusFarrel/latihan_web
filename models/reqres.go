package models

import (
	"time"

	"octaviusfarrel.dev/latihan_web/requests"
)

type ReqresUser struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type ReqresPostUser struct {
	ID string `json:"id"`
	requests.ReqresUserRequest
	CreatedAt time.Time `json:"createdAt"`
}
