package api

import (
	"encoding/json"
	"net/http"
)

type Sayhi struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func ReadHandler(response http.ResponseWriter, request *http.Request) {
	sayhi := Sayhi{Id: 1, Description: "hello world"}

	responseSayhi, err := json.Marshal(sayhi)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
	response.Write(responseSayhi)
}
