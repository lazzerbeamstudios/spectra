package home_api

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func Register(api huma.API) {

	huma.Register(api, huma.Operation{
		OperationID: "HomeGet",
		Method:      http.MethodGet,
		Path:        "/",
		Tags:        []string{"home"},
	}, HomeGetAPI)

}
