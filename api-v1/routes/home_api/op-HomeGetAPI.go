package home_api

import (
	"context"
)

type HomeGetOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

func HomeGetAPI(ctx context.Context, input *struct{}) (*HomeGetOutput, error) {
	response := &HomeGetOutput{}
	response.Body.Message = "Welcome to Lazzer Beam"
	return response, nil
}
