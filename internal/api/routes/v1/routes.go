package v1

import (
	"net/http"

	"github.com/basel2053/social-media/internal/api/handlers"
	"github.com/basel2053/social-media/internal/api/routes"
)

func RegisterV1Routes(mux *http.ServeMux, h *handlers.Handler) {
	group := routes.RouteGroup{
		Mux:    mux,
		Prefix: "/api/v1",
	}

	RegisterAuthRoutes(&group, h)
}
