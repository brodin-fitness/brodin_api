package app

import "net/http"

func StartApp() {
	r := initRouter()
	http.ListenAndServe(":8000", r)
}