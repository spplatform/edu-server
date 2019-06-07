package server

import (
	"fmt"
	"log"
	"net/http"
)

// RequestHandler is a main HTTP request handler
type RequestHandler struct {
}

// NewRequestHandler creates a new request handler with given set of dependencies
func NewRequestHandler() *RequestHandler {
	return &RequestHandler{}
}

// HandleHello handles hello requests
func (rh *RequestHandler) HandleHello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello request")
	fmt.Fprint(w, "Hello")
	w.WriteHeader(http.StatusOK)
}
