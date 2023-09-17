package main

import (
	"net/http"
)

func (a *Application) routes() *chi.Mux {
	a.use(a.Middleware.CheckRemember)

	a.get("/", a.Handlers.Home)

	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
