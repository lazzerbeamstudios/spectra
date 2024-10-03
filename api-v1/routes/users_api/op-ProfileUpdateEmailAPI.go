package users_api

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jinzhu/copier"

	"api-go/ent/user"
	"api-go/models"
	"api-go/utils/auth"
	"api-go/utils/db"
	"api-go/utils/encode"
)

type ProfileUpdateEmailInput struct {
	Auth string `header:"Authorization"`
	Body struct {
		Object models.ProfileUpdateEmail `json:"object"`
	}
}

func (input *ProfileUpdateEmailInput) Resolve(ctx huma.Context) []error {
	authID, authValid := auth.GetJWT(input.Auth)
	if !authValid {
		return []error{huma.Error401Unauthorized("Unable to authenticate.")}
	}
	if authID != input.Body.Object.ID {
		return []error{huma.Error401Unauthorized("User doesn't have permission.")}
	}

	profileObj, err := db.EntDB.User.Query().Where(user.ID(input.Body.Object.ID)).Only(context.Background())
	if err != nil {
		return []error{huma.Error404NotFound("User not found.")}
	}

	input.Body.Object.Email = strings.ToLower(input.Body.Object.Email)
	if input.Body.Object.Email == profileObj.Email {
		return []error{huma.Error400BadRequest("This email is not available.")}
	} else {
		usersEmail, err := db.EntDB.User.Query().Where(user.EmailEqualFold(input.Body.Object.Email)).All(context.Background())
		if len(usersEmail) > 0 || err != nil {
			return []error{huma.Error400BadRequest("This email is not available.")}
		}
	}

	return nil
}

type ProfileUpdateEmailOutput struct {
	Body struct {
		Object models.Profile `json:"object"`
	}
}

func ProfileUpdateEmailAPI(ctx context.Context, input *ProfileUpdateEmailInput) (*ProfileUpdateEmailOutput, error) {
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

	response := &ProfileUpdateEmailOutput{}
	object := &models.Profile{}
	copier.Copy(&object, &profileObj)
	response.Body.Object = *object
	return response, nil
}
