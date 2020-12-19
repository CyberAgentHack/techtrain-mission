package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/task4233/techtrain-mission/gameapi/handler"
	"github.com/task4233/techtrain-mission/gameapi/infra"
	"github.com/task4233/techtrain-mission/gameapi/usecase"
)

func main() {
	var port string = os.Getenv("PORT")
	srv := &http.Server{Addr: ":" + port}

	db, err := infra.NewDB()
	if err != nil {
		log.Println("failed newDB", err)
		os.Exit(1)
	}
	if db == nil {
		log.Println("failed newDB", err)
		os.Exit(1)
	}
	defer func() {
		cerr := db.Close()
		if err != nil {
			log.Println(cerr)
		}
	}()
	userRepo := infra.NewUserRepository(db)
	userUC := usecase.NewUser(userRepo)
	user := handler.NewUser(userUC)
	http.HandleFunc("/user/create", user.Create)
	http.HandleFunc("/user/get", user.Get)
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
