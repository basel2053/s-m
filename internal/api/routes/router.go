package routes

import (
	"fmt"
	"net/http"
)

type RouteGroup struct {
	Mux    *http.ServeMux
	Prefix string
}

func (rg *RouteGroup) HandleFunc(method, path string, handlerFunc http.HandlerFunc) {
	rg.Mux.HandleFunc(fmt.Sprintf("%s %s%s", method, rg.Prefix, path), handlerFunc)
}
