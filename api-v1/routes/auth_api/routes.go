package auth_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "Sign Up",
		Method:      http.MethodPost,
		Path:        "/auth/signup",
		Tags:        []string{"auth"},
	}, SignUpAPI)

	huma.Register(api, huma.Operation{
		OperationID: "Sign In",
		Method:      http.MethodPost,
		Path:        "/auth/signin",
		Tags:        []string{"auth"},
	}, SignInAPI)

}
