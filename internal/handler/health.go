package handler

import (
	"encoding/json"
	"net/http"
)

var ok = []byte("OK")

// HealthHandler type
type HealthHandler struct{}

// Health constructs a new health handler
func Health() HealthHandler {
	return HealthHandler{}
}

func (health HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		respond(w, errorJSON("only GET requests are supported"), http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write(ok)
}

func respond(w http.ResponseWriter, body []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = w.Write(body)
}

func errorJSON(msg string) []byte {
	s := struct {
		Error string `json:"error"`
	}{msg}
	bytes, _ := json.Marshal(s)
	return bytes
}
