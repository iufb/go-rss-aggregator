package main

import "net/http"

type status struct {
	status string
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, status{status: "NICE"})
}
