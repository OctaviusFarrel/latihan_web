package responses

import "octaviusfarrel.dev/latihan_web/models"

type (
	AllReqresUsers struct {
		BaseResponse
		ReqresUsers []models.ReqresUser `json:"users"`
	}

	ReqresUser struct {
		BaseResponse
		ReqresUser models.ReqresUser `json:"user"`
	}

	ReqresPostUser struct {
		BaseResponse
		ReqresPostUser models.ReqresPostUser `json:"user"`
	}
)
