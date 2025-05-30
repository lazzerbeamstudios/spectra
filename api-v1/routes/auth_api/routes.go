package auth_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "SignUpAPI",
		Method:      http.MethodPost,
		Path:        "/auth/signup",
		Tags:        []string{"auth"},
	}, SignUpAPI)

	huma.Register(api, huma.Operation{
		OperationID: "SignInAPI",
		Method:      http.MethodPost,
		Path:        "/auth/signin",
		Tags:        []string{"auth"},
	}, SignInAPI)

}
