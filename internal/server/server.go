package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

// Serve starts a server on the host and port specified
func Serve(host, port string) {
	r := mux.NewRouter()
	h := NewRequestHandler()
	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: r,
	}

	r.HandleFunc("/", h.HandleHello).Methods("GET")

	r.HandleFunc("/api/login", h.Login).Methods("GET", "POST")

	r.HandleFunc("/api/user/", h.GetUser).Methods("GET")
	// r.HandleFunc("/api/user", h.).Methods("POST")

	// r.HandleFunc("/api/poll", h.).Methods("GET")
	r.HandleFunc("/api/answer_poll", h.ProcessPolls).Methods("POST")

	// r.HandleFunc("/api/roadmap", h.).Methods("GET")
	// r.HandleFunc("/api/milestone", h.).Methods("GET")
	// r.HandleFunc("/api/step", h.).Methods("GET")
	// r.HandleFunc("/api/", h.).Methods("GET")
	// r.HandleFunc("/api/", h.).Methods("GET")
	// r.HandleFunc("/api/", h.).Methods("GET")
	// r.HandleFunc("/api/", h.).Methods("GET")

	http.Handle(host+":"+port, r)

	// handle interrupt
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-signalChan
		srv.Shutdown(context.Background())
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
