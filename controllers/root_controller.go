package controllers

import (
	"net/http"

	"github.com/brodin_fitness/brodin_api/utils"
)

func Ping(resp http.ResponseWriter, req *http.Request) {
	utils.Respond(resp, http.StatusOK, nil, "OK")
}
