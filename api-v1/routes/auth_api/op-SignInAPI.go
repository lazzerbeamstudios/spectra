package auth_api

import (
	"context"
	"strconv"

	"github.com/danielgtaylor/huma/v2"

	"api-go/ent/user"
	"api-go/utils/auth"
	"api-go/utils/db"
	"api-go/utils/password"
)

type SignInInput struct {
	Body struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
}

type SignInOutput struct {
	Body struct {
		Token string `json:"token"`
	}
}

func SignInAPI(ctx context.Context, input *SignInInput) (*SignInOutput, error) {
	userObj, err := db.EntDB.User.Query().Where(user.EmailEqualFold(input.Body.Email)).Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("User not found.")
	}

	passwordBool := password.CheckPassword(input.Body.Password, userObj.Password)
	if !passwordBool {
		return nil, huma.Error400BadRequest("Your password is incorrect.")
	}

	token, err := auth.CreateJWT(strconv.Itoa(userObj.ID), userObj.Email)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to authenticate.")
	}

	response := &SignInOutput{}
	response.Body.Token = token
	return response, nil
}
