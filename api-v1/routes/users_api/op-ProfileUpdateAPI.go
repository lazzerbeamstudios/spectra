package users_api

import (
	"context"
	"encoding/json"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/ent/user"
	"api-go/models"
	"api-go/utils/auth"
	"api-go/utils/db"
	"api-go/utils/encode"
)

type ProfileUpdateInput struct {
	Auth string `header:"Authorization"`
	Body struct {
		Object models.ProfileUpdate `json:"object"`
	}
}

func (input *ProfileUpdateInput) Resolve(ctx huma.Context) []error {
	authID, authValid := auth.GetJWT(input.Auth)
	if !authValid {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}
	if authID != input.Body.Object.ID {
		return []error{huma.Error401Unauthorized("User doesn't have permission.")}
	}
	return nil
}

type ProfileUpdateOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileUpdateAPI(ctx context.Context, input *ProfileUpdateInput) (*ProfileUpdateOutput, error) {
	userJson, _ := json.Marshal(input.Body.Object)
	userMap := encode.JsonToMap(string(userJson))
	delete(userMap, "id")

	_, err := db.BunDB.NewUpdate().Model(&userMap).TableExpr("users").Where("id=?", input.Body.Object.ID).Exec(ctx)
	if err != nil {
		return nil, huma.Error400BadRequest("Unable to update profile.")
	}

	profileObj, err := db.EntDB.User.Query().Where(user.ID(input.Body.Object.ID)).Only(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("User not found.")
	}

	response := &ProfileUpdateOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &profileObj)
	response.Body.Object = *object
	return response, nil
}
