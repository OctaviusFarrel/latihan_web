package responses

import (
	"net/http"

	"octaviusfarrel.dev/latihan_web/contants"
)

type BaseResponse struct {
	ResponseCode    string `json:"response_code"`
	ResponseMessage string `json:"response_message"`
}

func NewBaseResponse(code string, response *BaseResponse, err error) {
	switch code {
	case contants.CODE_200:
		response.ResponseCode = contants.CODE_200
		response.ResponseMessage = contants.MESSAGE_200
	case contants.CODE_400:
		response.ResponseCode = contants.CODE_400
		response.ResponseMessage = contants.MESSAGE_400
	case contants.CODE_401:
		response.ResponseCode = contants.CODE_401
		response.ResponseMessage = contants.MESSAGE_401
	case contants.CODE_404:
		response.ResponseCode = contants.CODE_404
		response.ResponseMessage = contants.MESSAGE_404
	case contants.CODE_500:
		response.ResponseCode = contants.CODE_500
		response.ResponseMessage = contants.MESSAGE_500
	default:
		response.ResponseCode = contants.CODE_500
		response.ResponseMessage = contants.MESSAGE_500
	}
}

func NewBaseResponseStatusCode(code int, response *BaseResponse, err error) {
	switch code {
	case http.StatusOK:
		response.ResponseCode = contants.CODE_200
		response.ResponseMessage = contants.MESSAGE_200
	case http.StatusBadRequest:
		response.ResponseCode = contants.CODE_400
		response.ResponseMessage = contants.MESSAGE_400
	case http.StatusBadGateway:
		response.ResponseCode = contants.CODE_401
		response.ResponseMessage = contants.MESSAGE_401
	case http.StatusNotFound:
		response.ResponseCode = contants.CODE_404
		response.ResponseMessage = contants.MESSAGE_404
	case http.StatusInternalServerError:
		response.ResponseCode = contants.CODE_500
		response.ResponseMessage = contants.MESSAGE_500
	default:
		response.ResponseCode = contants.CODE_500
		response.ResponseMessage = contants.MESSAGE_500
	}
}
