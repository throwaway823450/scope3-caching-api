package api

import (
	"net/http"
	"scope3/caching-api/measurement"

	"github.com/gorilla/mux"
)

func NewRouter(measurementClient measurement.CachingClient) http.Handler {
	h := NewHandler(measurementClient)
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/emissions", h.PostEmmisions).Methods("POST")

	return r
}
