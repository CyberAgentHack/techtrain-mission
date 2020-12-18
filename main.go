package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// TestHandler is for testing
func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %s\n", r.URL)

	var connString string = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	) + "?parseTime=true&collation=utf8mb4_bin"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed sql.Open: %s", err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT version()")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed Query: %s", err)
		return
	}
	if ok := rows.Next(); !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed rows.Next"))
		return
	}

	var version string
	if err = rows.Scan(&version); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "failed Scan: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "mysql version: %s\n", version)
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
