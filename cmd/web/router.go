package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/ping", app.ping)
	router.HandlerFunc(http.MethodGet, "/counter", app.counter)
	router.HandlerFunc(http.MethodGet, "/insert/:name", app.insert)
	router.HandlerFunc(http.MethodGet, "/latest", app.latest)

	return router
}
