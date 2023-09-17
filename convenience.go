package main

import "net/http"

func (a *Application) get(s string, h http.HandlerFunc) {
	a.App.Routes.Get(s, h)
}

func (a *Application) post(s string, h http.HandlerFunc) {
	a.App.Routes.Post(s, h)
}

func (a *Application) use(m ...func(handler http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}
