package app

import (
	"log"
	"net/http"

	"github.com/brodin_fitness/brodin_api/config"
)

// Initialize env variables and routes, then start the server.
func Start() {
	initEnv()
	config.Values = config.Init()
	r := initRouter()

	host := ":" + config.Values.App.Port
	log.Println("🚀 App listening on port", config.Values.App.Port + " 🚀")
	http.ListenAndServe(host, r)
}
