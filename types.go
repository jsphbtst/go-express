package express

import "net/http"

type Config map[string]string

// just so there's semantic difference for the reader
type Handler = func(http.ResponseWriter, *http.Request)

type Server struct {
	*http.Server
}

type GenericResponse = map[string]string

type Route = map[string]Handler

type Express struct {
	routes      StringSet
	middlewares []Handler

	getRoutes   StringSet
	getHandlers Route

	postRoutes   StringSet
	postHandlers Route

	putRoutes   StringSet
	putHandlers Route

	patchRoutes   StringSet
	patchHandlers Route

	deleteRoutes   StringSet
	deleteHandlers Route
}
