package express

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	*http.Server
}

type GenericResponse map[string]string

// just so there's semantic difference for the reader
type Handler = func(http.ResponseWriter, *http.Request)

type Route map[string]Handler

type Express struct {
	routes       StringSet
	middlewares  []Handler
	getRoutes    StringSet
	getHandlers  Route
	postRoutes   StringSet
	postHandlers Route
}

func response404Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// note to self: this goes strictly AFTER else it does not
	// get parsed as JSON lul
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(GenericResponse{"message": "Error"})
}

func New() *Express {
	return &Express{
		routes:       []string{},
		middlewares:  [](func(http.ResponseWriter, *http.Request)){},
		getRoutes:    []string{},
		postRoutes:   []string{},
		getHandlers:  make(Route),
		postHandlers: make(Route),
	}
}

func (app *Express) Use(middleware Handler) {
	app.middlewares = append(app.middlewares, middleware)
}

func (app *Express) Get(pathname string, handler Handler) {
	app.getHandlers[pathname] = handler
	app.routes.Add(pathname)
	app.getRoutes.Add(pathname)
}

func (app *Express) Post(pathname string, handler Handler) {
	app.postHandlers[pathname] = handler
	app.routes.Add(pathname)
	app.postRoutes.Add(pathname)
}

func (app *Express) Listen(port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range app.middlewares {
			middleware(w, r)
		}

		currentPath := r.URL.Path
		isPathExists := app.routes.Contains(currentPath)
		if !isPathExists {
			response404Handler(w, r)
			return
		}

		if r.Method == http.MethodGet && app.getRoutes.Contains(currentPath) {
			getHandler := app.getHandlers[currentPath]
			getHandler(w, r)
			return
		}

		if r.Method == http.MethodPost && app.postRoutes.Contains(currentPath) {
			postHandler := app.postHandlers[currentPath]
			postHandler(w, r)
			return
		}

		response404Handler(w, r)
	})

	PORT := fmt.Sprintf(":%d", port)
	log.Printf("HTTP Server listening on PORT %s\n", PORT)
	http.ListenAndServe(PORT, nil)
}
