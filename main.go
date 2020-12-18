package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TestHandler is for testing
func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s\n", r.URL)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("test handler"))
	return
}

func main() {
	var port string = os.Getenv("PORT")
	srv := &http.Server{Addr: ":" + port}
	http.HandleFunc("/", TestHandler)
	log.Printf("Start App: listening on port %s", port)

	// graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("shutdown the server with error: %+v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("failed to shutdown: %+v", err)
		os.Exit(1)
	}
}
