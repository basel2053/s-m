package v1

import (
	"github.com/basel2053/social-media/internal/api/handlers"
	"github.com/basel2053/social-media/internal/api/routes"
)

func RegisterAuthRoutes(rg *routes.RouteGroup, h *handlers.Handler) {
	rg.HandleFunc("POST", "/signup", h.Signup)
}
