package app

import (
	"net/http"

	"github.com/brodin_fitness/brodin_api/config"
)

// Initialize env variables and routes, then start the server.
func Start() {
	initEnv()
	config.Values.Init()
	r := initRouter()

	http.ListenAndServe(config.Values.App.Port, r)
}
