package home_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "HomeGetAPI",
		Method:      http.MethodGet,
		Path:        "/",
		Tags:        []string{"home"},
	}, HomeGetAPI)

}
