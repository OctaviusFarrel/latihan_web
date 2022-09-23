package external

import (
	"testing"

	"octaviusfarrel.dev/latihan_web/models"
)

func TestReqresExternal_AllUsers(t *testing.T) {
	type result struct {
		Data []models.ReqresUser `json:"data"`
	}
	tests := []struct {
		name  string
		wants result
	}{
		{wants: result{
			Data: []models.ReqresUser{
				{
					Id:        1,
					Email:     "george.bluth@reqres.in",
					FirstName: "George",
					LastName:  "Bluth",
					Avatar:    "https://reqres.in/img/faces/1-image.jpg",
				},
			},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testR := result{}
			r := ReqresExternal{}

			r.AllUsers(&testR)

			if testR.Data[0] != tt.wants.Data[0] {
				t.Errorf("Result: %q, must've been : %q", testR.Data[0], tt.wants.Data[0])
			}
		})
	}
}
