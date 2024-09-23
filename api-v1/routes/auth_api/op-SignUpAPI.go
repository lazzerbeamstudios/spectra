package auth_api

import (
	"context"
	"strconv"
	"strings"

	"github.com/danielgtaylor/huma/v2"

	"api-go/ent/user"
	"api-go/utils/auth"
	"api-go/utils/db"
	"api-go/utils/password"
)

type SignUpInput struct {
	Body struct {
		Password string `json:"password" minLength:"8" maxLength:"30"`
		Email    string `json:"email" minLength:"2" maxLength:"30" format:"email"`
	}
}

func (input *SignUpInput) Resolve(ctx huma.Context) []error {
	input.Body.Email = strings.ToLower(input.Body.Email)
	usersList, err := db.EntDB.User.Query().Where(user.EmailEqualFold(input.Body.Email)).All(context.Background())
	if len(usersList) > 0 || err != nil {
		return []error{huma.Error400BadRequest("This email is not available.")}
	}

	input.Body.Password, err = password.HashPassword(input.Body.Password)
	if err != nil {
		return []error{huma.Error400BadRequest("Unable to create user.")}
	}

	return nil
}

type SignUpOutput struct {
	Body struct {
		Token string `json:"token"`
	}
}

func SignUpAPI(ctx context.Context, input *SignUpInput) (*SignUpOutput, error) {
	userObj, err := db.EntDB.User.Create().
		SetPassword(input.Body.Password).
		SetEmail(input.Body.Email).
		Save(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to create user.")
	}

	token, err := auth.CreateJWT(strconv.Itoa(userObj.ID), userObj.Email)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to authenticate.")
	}

	response := &SignUpOutput{}
	response.Body.Token = token
	return response, nil
}
