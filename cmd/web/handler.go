package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("Ping Hello Selam")
	w.Write([]byte("Hello Selam"))
}

func (app *application) counter(w http.ResponseWriter, r *http.Request) {
	app.logger.Info("increment counter")
	val, err := app.redisClient.IncrBy(r.Context(), "counter", 1).Result()
	if err != nil {
		app.serverError(w, r, err)
	}
	counterString := fmt.Sprintf("You have visited this website %d times", val)
	w.Write([]byte(counterString))
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	name := params.ByName("name")
	greetings := fmt.Sprintf("Hello there %s", name)
	err := app.user.Insert(name)
	if err != nil {
		app.serverError(w, r, err)
	}
	w.Write([]byte(greetings))
}

func (app *application) latest(w http.ResponseWriter, r *http.Request) {
	users, err := app.user.Latest()
	if err != nil {
		app.serverError(w, r, err)
	}
	latestUsers, err := json.Marshal(users)
	if err != nil {
		app.serverError(w, r, err)
	}
	app.logger.Info("Displaying latest...", "latest", users)
	w.Write(latestUsers)
}
