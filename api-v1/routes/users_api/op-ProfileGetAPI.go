package users_api

import (
	"context"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/ent/user"
	"api-go/models"
	"api-go/utils/auth"
	"api-go/utils/db"
)

type ProfileGetInput struct {
	Auth string `header:"Authorization"`
}

type ProfileGetOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileGetAPI(ctx context.Context, input *ProfileGetInput) (*ProfileGetOutput, error) {
	authID, authValid := auth.GetJWT(input.Auth)
	if !authValid {
		return nil, huma.Error401Unauthorized("Unable to authenticate.")
	}

	profileObj, err := db.EntDB.User.Query().Where(user.ID(authID)).Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("User not found.")
	}

	response := &ProfileGetOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &profileObj)
	response.Body.Object = *object
	return response, nil
}
