package utils

import (
	"encoding/json"
	"net/http"
)

type body struct {
	Status string `json:"status"`
	Message string `json:"msg"`
	Data interface{} `json:"data"`
}

// Convert the provided data into JSON then send it to the user
func Respond(resp http.ResponseWriter, status int, data interface{}, msg string) {
	body, _ := json.Marshal(body{
		Status: "success",
		Message: msg,
		Data: data,
	})

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	resp.Write(body)
}