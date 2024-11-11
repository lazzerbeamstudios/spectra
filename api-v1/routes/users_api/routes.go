package users_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"

	"api-go/middleware"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "Profile Get",
		Method:      http.MethodGet,
		Path:        "/users/profile",
		Tags:        []string{"users"},
		Middlewares: huma.Middlewares{middleware.AuthMiddleware},
	}, ProfileGetAPI)

	huma.Register(api, huma.Operation{
		OperationID: "Profile Update",
		Method:      http.MethodPut,
		Path:        "/users/profile",
		Tags:        []string{"users"},
	}, ProfileUpdateAPI)

	huma.Register(api, huma.Operation{
		OperationID: "Profile Update Email",
		Method:      http.MethodPut,
		Path:        "/users/profile/email",
		Tags:        []string{"users"},
	}, ProfileUpdateEmailAPI)

}
