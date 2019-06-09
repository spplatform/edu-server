package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Serve starts a server on the host and port specified
func Serve(host, port string) {
	lm, err := NewLogicManager()
	if err != nil {
		log.Fatal(err)
	}
	h := NewRequestHandler(lm)

	r := mux.NewRouter()
	rh := handlers.RecoveryHandler()(r)
	srv := &http.Server{
		Addr:    host + ":" + port,
		Handler: rh,
	}

	r.HandleFunc("/health", h.HandleHello).Methods("GET")

	r.HandleFunc("/api/login", h.Login).Methods("GET", "POST")

	r.HandleFunc("/api/user/{id:[0-9]+}", h.GetUser).Methods("GET")
	r.HandleFunc("/api/user/{id:[0-9]+}/process-poll", h.ProcessPolls).Methods("POST")

	r.HandleFunc("/api/roadmap/{id:[0-9]+}", h.GetRoadmap).Methods("GET")

	r.HandleFunc("/api/badge/{id:[0-9]+}", h.GetBadge).Methods("GET")
	r.HandleFunc("/api/badge", h.IssueBadge).Methods("POST")

	r.HandleFunc("/api/certificate/{id:[0-9]+}", h.GetCertificate).Methods("GET")
	r.HandleFunc("/api/certificate", h.IssueCertificate).Methods("POST")

	http.Handle(host+":"+port, r)

	// handle interrupt
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-signalChan
		srv.Shutdown(context.Background())
	}()

	// start server
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
