package httpapi

import (
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthcheck)
	mux.HandleFunc("GET /", index)

	return mux
}

func healthcheck(writer http.ResponseWriter, request *http.Request) {
	writeJSON(writer, http.StatusOK, response{
		Message:   "vectora server is healthy",
		Status:    "ok",
		Timestamp: time.Now().UTC(),
	})
}

func index(writer http.ResponseWriter, request *http.Request) {
	writeJSON(writer, http.StatusOK, response{
		Message:   "vectora server is running",
		Status:    "ok",
		Timestamp: time.Now().UTC(),
	})
}

func writeJSON(writer http.ResponseWriter, statusCode int, payload response) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(payload); err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
