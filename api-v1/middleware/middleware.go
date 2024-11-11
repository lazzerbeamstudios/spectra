package middleware

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

func AuthMiddleware(ctx huma.Context, next func(huma.Context)) {

	fmt.Println("AuthMiddleware")

}
