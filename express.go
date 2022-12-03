package express

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func response404Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// note to self: this goes strictly AFTER else it does not
	// get parsed as JSON lul
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(GenericResponse{"message": "Error"})
}

func New() *Express {
	return &Express{
		routes:         []string{},
		middlewares:    [](func(http.ResponseWriter, *http.Request)){},
		getRoutes:      []string{},
		postRoutes:     []string{},
		getHandlers:    make(Route),
		postHandlers:   make(Route),
		putRoutes:      []string{},
		putHandlers:    make(Route),
		patchRoutes:    []string{},
		patchHandlers:  make(Route),
		deleteRoutes:   []string{},
		deleteHandlers: make(Route),
	}
}

func (app *Express) Use(handler Handler) {
	app.middlewares = append(app.middlewares, handler)
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

func (app *Express) Put(pathname string, handler Handler) {
	app.putHandlers[pathname] = handler
	app.routes.Add(pathname)
	app.putRoutes.Add(pathname)
}

func (app *Express) Patch(pathname string, handler Handler) {
	app.patchHandlers[pathname] = handler
	app.routes.Add(pathname)
	app.patchRoutes.Add(pathname)
}

func (app *Express) Delete(pathname string, handler Handler) {
	app.deleteHandlers[pathname] = handler
	app.routes.Add(pathname)
	app.deleteRoutes.Add(pathname)
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

		if r.Method == http.MethodPut && app.putRoutes.Contains(currentPath) {
			putHandler := app.putHandlers[currentPath]
			putHandler(w, r)
			return
		}

		if r.Method == http.MethodPatch && app.patchRoutes.Contains(currentPath) {
			patchHandler := app.patchHandlers[currentPath]
			patchHandler(w, r)
			return
		}

		// DELETE
		if r.Method == http.MethodDelete && app.deleteRoutes.Contains(currentPath) {
			deleteHandler := app.deleteHandlers[currentPath]
			deleteHandler(w, r)
			return
		}

		response404Handler(w, r)
	})

	PORT := fmt.Sprintf(":%d", port)
	log.Printf("HTTP Server listening on PORT %s\n", PORT)
	http.ListenAndServe(PORT, nil)
}
