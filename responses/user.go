package responses

import "octaviusfarrel.dev/latihan_web/models"

type (
	User struct {
		BaseResponse
		User models.UserModel `json:"user"`
	}

	UserWithToken struct {
		User
		Token string `json:"token"`
	}
)
