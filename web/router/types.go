package router

import (
	"net/http"
)

// Router is the object we can use to connect our
// http.Handlers to the application's router
type Router interface {
	Handle(path string, handler http.Handler)
}
