package app

import "net/http"

// Initialize env variables and routes, then start the server.
func StartApp() {
	initEnv()
	r := initRouter()
	http.ListenAndServe(":8000", r)
}